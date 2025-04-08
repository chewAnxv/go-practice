package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 3
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)
	//len返回‘对数’
	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(m)

	// 能删除所有的函数
	clear(m)
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println(prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n2 === n")
	}


}
