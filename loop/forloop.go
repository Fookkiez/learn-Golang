package main

import "fmt"

var score int

func main() {

	i := 0
	for i < 10 {
		fmt.Println("i = ", i)
		i = i + 1
	}

	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input", input)
		if input == "exit" {
			break
		}
	}

}
