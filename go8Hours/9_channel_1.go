package main

import "fmt"

func main() {
	// 创建一个channel管道，用来装入类型为int的数值
	a := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine正在运行...")
		a <- 666 // 将666这个值塞进管道里面
	}()

	num := <-a // 从a中接受数据，并赋值给num
	fmt.Println("num = ", num)
}
