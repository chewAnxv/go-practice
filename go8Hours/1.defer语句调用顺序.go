package main

import "fmt"

func deferFunc() int {
	fmt.Println("Defer first")

	return 0
}

func returnFunc() int {
	fmt.Println("Return first")

	return 0
}

func deferReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	fmt.Println(deferReturn())
}
