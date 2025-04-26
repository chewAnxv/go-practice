package main

import "fmt"

func main() {
	c := make(chan int)

	// 子goroutine
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("塞入一个channel")
		}

		// close()关键字用来关闭c通道
		close(c)
	}()

	for {
		// data接受值，
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main finished.")

}
