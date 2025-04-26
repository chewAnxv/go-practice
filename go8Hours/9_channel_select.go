package main

import "fmt"

func fibonacii(a, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case a <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	a := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println(<-a)
		}

		quit <- 0

	}()

	// main go
	fibonacii(a, quit)
}
