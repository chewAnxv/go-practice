package main

import "fmt"

type Man struct {
	name string
	age  int
}

func (this *Man) Walk() {
	fmt.Println("Man can Walk")
}

func (this *Man) Eat() {
	fmt.Println("Man can Eat")
}

// 创建一个SuperMan，SuperMan能继承man的属性

type SuperMan struct {
	Man
	level int
}

func (this SuperMan) Show() {
	fmt.Println("Name:", this.name)
	fmt.Println("Age", this.age)
	fmt.Println("Level:", this.level, "\n")
}

func (this SuperMan) Fly() {
	fmt.Println("SuperMan can fly")
}

func main() {
	human := Man{name: "ZhangSan", age: 10}
	human.Walk()
	human.Eat()


	fmt.Println("--------")


	var superHuman SuperMan
	superHuman.name = "LiSi"
	superHuman.age = 20
	superHuman.level = 1
	// 虽然继承了Man，调用的是Man的方法的化还是原来的方法
	superHuman.Show()
	superHuman.Eat()
	superHuman.Fly()
}
