package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {

	// Json to Object ==> eg.  e.ID -> GET VALUES
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"sds","Tel":"01101","Email":"sdsd@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
}
