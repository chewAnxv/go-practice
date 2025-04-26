package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("defer A")

		func() {
			defer fmt.Println("defer B")

			// 能直接退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		// 匿名函数结尾的空（）意思是调用这个函数

		fmt.Println("A")
	}()
	// 匿名函数结尾的空（）意思是调用这个函数


	
	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
