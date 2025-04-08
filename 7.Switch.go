package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Is before noon")
	default:
		fmt.Println("Is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("I am string.")
		case int:
			fmt.Println("I am int.")
		case bool:
			fmt.Println("I am bool.")
		default:
			fmt.Println("Don't know the type %T\n", t)
		}
	}
	whatAmI("你好")
	whatAmI(true)
	whatAmI(111)
}
