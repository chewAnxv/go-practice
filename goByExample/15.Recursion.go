package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
	//每次都在调用函数，直到函数表示出来了真正的值之后，再用值来反推答案进行计算
	fmt.Println(fib(6) + fib(5))
	fmt.Println(fib(5) + fib(4) + fib(4) + fib(3))
	fmt.Println(fib(4) + fib(3) + fib(3) + fib(2) + fib(3) + fib(2) + fib(2) + fib(1))
	fmt.Println(fib(3) + fib(2) + fib(2) + fib(1) + fib(2) + fib(1) + fib(2) + fib(2) + fib(1) + fib(1) + fib(0) + fib(1) + fib(0) + fib(1))
}
