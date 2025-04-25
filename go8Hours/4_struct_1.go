package main

import "fmt"

type book struct {
	auth  string
	title string
}

func changeBook1(b *book) {
	// 有指针，需要用&传递参数的地址
	b.auth = "MoYan"
	b.title = "buzhidao"
}

func changeBook2(b book) {
	// 没有指针，传递的是参数的副本
	b.auth = "MoYan"
	b.title = "buzhidao"

}

func main() {
	// 声明一个book1变量，类型是book
	b.title = "buzhidao"
	var book1 book

	book1.auth = "YuHua"
	book1.title = "Huozhe"

	fmt.Println(book1)

	changeBook1(&book1)
	fmt.Println(book1)

	changeBook2(book1)
	fmt.Println(book1)
}
