package main

import "fmt"

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3val(val int, val2 int, val3 int) int {
	return val + val2 + val3
}

func main() {
	result := plus(5, 2)
	fmt.Println("result", result)
	result = plus3val(1, 2, 3)
	fmt.Println("result", result)
}
