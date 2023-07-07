package main

import "fmt"

var score int

func main() {

	fmt.Scanf("%d", &score)
	if score >= 80 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else if score >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	i := 0
	for i < 10 {
		fmt.Println("i = ", i)
		i = i + 1
	}

}
