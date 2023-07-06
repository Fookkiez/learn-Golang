package main

import (
	"fmt"
)

var product = make(map[string]float64)

func main() {
	fmt.Println(product)

	// add
	product["Mac"] = 12
	product["PC"] = 13
	product["Ipad"] = 14
	fmt.Println(product)

	//delete
	delete(product, "PC")
	fmt.Println(product)

	//Update
	product["Ipad"] = 33
	fmt.Println(product)

	value1 := product["iphone"]
	fmt.Println(value1)

	courseName := map[string]string{"101":"Java","102":"C#","103":"Python"}
	fmt.Println(courseName)

}