package main

import "fmt"

var courseName []string

func main() {

	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)

	courseName = append(courseName, "C#", "C", "Html")
	fmt.Println(courseName)

	courseWeb := courseName[3:5]
	fmt.Println(courseWeb)

	courseWeb = courseName[1:3]
	fmt.Println(courseWeb)

}
