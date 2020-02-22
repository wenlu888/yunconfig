package main

import "yunConfig/allConfig"

func main() {
	path := `D:\工作\云桌面\福建资源池\福建资源池v1.1.xlsx`
	//allConfig.GetComputConfig(path)
	allConfig.GetCephNodeList(path)
}
