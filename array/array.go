package main

import (
	"fmt"
)

var productName [4]string
var price [4]float32

func main() {

	productName[0] = "s"
	productName[1] = "a"
	productName[2] = "c"
	productName[3] = "d"

	productName := [4]string{"zz", "aa", "vv", "cc"}
	price := [4]float32{21, 22, 23, 24}

	fmt.Println(productName)
	fmt.Println(price)
}
