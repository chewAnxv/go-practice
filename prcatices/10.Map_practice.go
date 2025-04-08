package main

import "fmt"

func main() {
	fruitPrices := make(map[string]int)

	fruitPrices["apple"] = 5
	fruitPrices["banana"] = 3
	fruitPrices["orange"] = 4
	fmt.Println(fruitPrices)

	values, sold := fruitPrices["mango"]
	fmt.Println(values, sold)
	_, exists := fruitPrices["mango"]
	if exists {
		fmt.Println("Mango available.")
	} else {
		fmt.Println("Mango is sold out.")
	}

}
