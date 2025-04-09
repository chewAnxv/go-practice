package main

import "fmt"

func calculate(x, y int) (int, int, int) {
	return x + y, x - y, x * y

}
func main() {
	sum, difference, product := calculate(5, 3)
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	sum, _, product = calculate(5, 3)
	fmt.Println(sum)
	fmt.Println(product)
}
