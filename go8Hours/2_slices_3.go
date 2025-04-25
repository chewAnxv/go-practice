package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// 	这里截取的是第0，1元素，没有
	s1 := s[0:2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	// copy可以将s的值拷贝到s2中去
	s2 := make([]int, 3)

	fmt.Println(s2)

	copy(s2, s)
	
	fmt.Println(s2)
}
