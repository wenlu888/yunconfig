package readExcel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)


type Comput struct {
	Title  		string 		// 节点类型
	MargeIp		string		// 管理ip  host_ip
	ServiceIp	string 		// 业务ip  spice_port_bind_ip
	IsManager   bool        // 是否为控制节点
}

// 读取计算节点
func ReadExcelComput(path string)(a []Comput,err error){
	file,err := excelize.OpenFile(path)
	if err != nil {return nil,err}
	data,err := file.GetRows("计算服务器")
	if err != nil {return nil,err}
	//fmt.Println(data)
	for i:=1;i<len(data);i++{
		var c Comput
		c.MargeIp = data[i][2]
		c.ServiceIp = data[i][4]
		if len(data[i]) > 8 {
			c.IsManager = true
		}
		c.Title = data[i][1]
		if !strings.Contains(data[i][1],"计算节点") {continue}
		a = append(a, c)
	}
	return a,nil
}

type CephNode struct {
	Public  string
	Cluster string
	Mon     bool
}

func ReadExcelCeph(path string)(a []CephNode,err error){
	file,err := excelize.OpenFile(path)
	if err != nil {return nil,err}
	cephData,err := file.GetRows("存储服务器")
	if err != nil {return nil,err}
	//fmt.Println(data)
	for i:=1;i<len(cephData);i++{
		var c CephNode
		c.Public = cephData[i][3]
		c.Cluster = cephData[i][4]
		if len(cephData[i]) > 6 {
			if strings.Contains(cephData[i][5],"MON"){
				c.Mon = true
			}else if strings.Contains(cephData[i][5],"管理") || strings.Contains(cephData[i][5],"预留") {
				continue
			}
		}
		a = append(a, c)
	}
	return a,nil
}