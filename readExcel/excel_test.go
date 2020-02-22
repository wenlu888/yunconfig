package readExcel

import (
	"fmt"
	"testing"
)

func TestExcel(t *testing.T){
	a,err := ReadExcelCeph(`D:\工作\云桌面\福建资源池\福建资源池v1.1.xlsx`)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a)

}
