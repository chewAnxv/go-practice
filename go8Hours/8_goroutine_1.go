package main

import (
	"fmt"
	"time"
)

// 从goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goutine : i = %d", i)
		time.Sleep(1 * time.Second)
	}
}


// 主goroutine
func main() {
	
	// 创建一个go程，去执行newTask任务
	go newTask()
	
	// main退出其它goroutine会消失
	fmt.Println("exist")

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("new Goutine : i = %d", i)
	// 	time.Sleep(1 * time.Second)
	// }
	
}
