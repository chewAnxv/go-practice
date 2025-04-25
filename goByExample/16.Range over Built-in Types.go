package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//只写一个变量go编译器默认只把字典里面的键"key"给你
	for k := range kvs {
		fmt.Println(k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
