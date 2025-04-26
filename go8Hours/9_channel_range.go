package main

import "fmt"

func main() {
	c := make(chan int)

	// 子goroutine
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close()关键字用来关闭c通道
		close(c)
	}()

	// range可以不断迭代channel操作
	for data := range c {
		fmt.Println(data)
	}


	fmt.Println("Main finished.")

}
