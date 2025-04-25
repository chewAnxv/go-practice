package main

import "fmt"

type AnimalIf interface {
	Sleep()
	Color() string
	GetType() string
}

// 猫类型
type Cat struct {
	color string
}

func (C *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (C *Cat) Color() string {
	return C.color
}
func (C *Cat) GetType() string {
	return "Cat"
}

// 狗类型
type Dog struct {
	color string
}

func (C *Dog) Sleep() {
	fmt.Println("dog is sleep")
}

func (C *Dog) Color() string {
	return C.color
}
func (C *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIf) {
	animal.Sleep()
	fmt.Println(animal.Color())
	fmt.Println(animal.GetType())
}

func main() {
	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

}
