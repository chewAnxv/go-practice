package main

import "fmt"

func main() {
	var s []string
	s = make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(len(s))

	s = append(s, "f")

	// copy()方法将s赋值给c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:2]
	fmt.Println(l)

	l = s[3:]
	fmt.Println(l)

	// [][]int：表示**“切片的切片”，
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			// 访问二维切片 twoD 中第 i 行、第 j 列的元素，并给它赋值为 i + j。
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
