package main

import "fmt"

func main() {
	//声明slice1是一个切片，同时给slices1分配3个空间，初始化的值为0
	//通过:=推导出slices1是一个切片
	slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//判断slice是否为空
	if slice1 == nil {
		fmt.Println("slices1 is nil")
	} else {
		fmt.Println("slices1 is not nil")
	}
}
