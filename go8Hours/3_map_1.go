package main

import "fmt"

func main() {
	// 声明一个myMap1，类型为map，键是string类型，值也是string类型
	myMap1 := make(map[string]string)

	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "go"

	// 因为是哈希值填入map，所以输出出来是乱序
	fmt.Println(myMap1)

	// 声明方式2
	myMap2 := map[int]string{
		1: "java",
		2: "php",
		3: "python",
	}

	fmt.Println(myMap2)

}
