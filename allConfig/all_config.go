package allConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"yunConfig/readExcel"
)

type openStack struct {
	DeployFlag   	 string 			 `json:"deploy_flag"`
	OsPassword   	 string 			 `json:"os_password"`
	OsRegionName  	 string 			 `json:"os_region_name"`
	FirstRegionOsVip string 			 `json:"first_region_os_vip"`
	OsUseNginx    	 string 			 `json:"os_use_nginx"`
	ControllerIp     []string 			 `json:"controller_ip"`
	ComputeNode      []computeNode 		 `json:"compute_node"`
	NetworkNode      []string            `json:"network_node"`
}

type computeNode struct {
	HostIp  		  string 	`json:"host_ip"`
	SpicePortBindIp   string 	`json:"spice_port_bind_ip"`
}

func GetComputConfig(path string){
	c,err := readExcel.ReadExcelComput(path)
	if err != nil {fmt.Println("read Excel error:"+err.Error())}
	var coList []computeNode
	for _,i := range c {
		if !i.IsManager {
			var co computeNode
			co.HostIp = i.MargeIp
			co.SpicePortBindIp = i.ServiceIp
			coList = append(coList,co)
		}
	}
	coJson,err := json.Marshal(&coList)
	if err != nil {
		fmt.Println("生成json字符串异常：",err.Error())
	}
	ioutil.WriteFile("computeNode.txt",coJson,0777)
}


// Ceph 节点
type CephNode struct{
	PublicAddr   string  `json:"public_addr"`
	ClusterAddr  string  `json:"cluster_addr"`
	IsMon        string  `json:"is_mon"`
	WalDbDevices []string `json:"wal_db_devices"`
	DataDevices  []string `json:"data_devices"`
}


func GetCephNodeList(path string){
	cephList,err := readExcel.ReadExcelCeph(path)
	if err != nil {fmt.Println("read ceph excel error :",err.Error())}
	var cephL []CephNode
	for _,c := range cephList {
		var ceph CephNode
		ceph.PublicAddr = c.Public
		ceph.ClusterAddr = c.Cluster
		ceph.IsMon = strconv.FormatBool(c.Mon)
		ceph.DataDevices = []string{
			"/dev/sdb",
			"/dev/sdc",
			"/dev/sdd",
			"/dev/sde",
			"/dev/sdf",
			"/dev/sdg",
			"/dev/sdh",
			"/dev/sdi",
			"/dev/sdj",
			"/dev/sdk",
			"/dev/sdl",
			"/dev/sdm",
		}
		ceph.WalDbDevices = []string{"/dev/nvme0n1p2","/dev/nvme1n1p2"}
		cephL = append(cephL,ceph)
	}
	cephJson,err := json.Marshal(&cephL)
	if err != nil {fmt.Println("序列化ceph Json失败",err.Error())}
	err = ioutil.WriteFile("cephNode.txt",cephJson,0777)
	if err != nil {
		fmt.Println("生成ceph json文件异常",err.Error())
	}
}

