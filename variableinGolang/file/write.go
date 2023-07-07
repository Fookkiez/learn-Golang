package main

import "os"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("hello\n FOokkie")
	err := os.WriteFile("data.txt", data1, 0644)
	Check(err)
	f, err := os.Create("employeeTest")
	Check(err)

	defer f.Close()

	data2 := []byte("Fookie \n Nantha")
	os.WriteFile("EmployeeCreatefile.txt", data2, 0644)

}
