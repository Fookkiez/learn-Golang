package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	//Json Marshal อ่านเฉพาะ Key ขึ้นต้นด้วยตัวใหญ่เท่านั้น
	data, _ := json.Marshal(employee{101, "sds", "01101", "sdsd@gmail.com"})
	fmt.Println(string(data))

}
