package main

import "fmt"

// "interface{}" 万能类型，能接受所有类型
func myFunc(arg interface{}) {
	fmt.Println("my Func is called:")
	fmt.Println(arg)

	// 类型断言
	value, ok := arg.(string)

	if !ok {
		fmt.Println("my Func is not string")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
	fmt.Println("--------")
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)

	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
	myFunc(true)

}
