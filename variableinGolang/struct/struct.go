package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeID:   "100",
		employeeName: "Pradoo",
		phone:        "090000000",
	}
	employeeList[1] = employee{
		employeeID:   "101",
		employeeName: "“Prayad”",
		phone:        "090000001",
	}
	employeeList[2] = employee{
		employeeID:   "102",
		employeeName: "“Pranee”",
		phone:        "090000002",
	}

	fmt.Println("employee = ", employeeList)

	employeeListslice := []employee{}

	employee0 := employee{
		employeeID:   "100",
		employeeName: "Pradoo",
		phone:        "090000000",
	}

	employee1 := employee{
		employeeID:   "101",
		employeeName: "Prayad",
		phone:        "090000001",
	}

	employee2 := employee{
		employeeID:   "102",
		employeeName: "Pranee",
		phone:        "090000002",
	}

	employeeListslice = append(employeeListslice, employee0)
	employeeListslice = append(employeeListslice, employee1)
	employeeListslice = append(employeeListslice, employee2)

	fmt.Println("employeeListslice = ", employeeListslice)

}
