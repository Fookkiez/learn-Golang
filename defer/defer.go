package main

import "fmt"

func add(va1, val2 float64) {
	result := va1 + val2
	fmt.Println("Result : ", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j = ", j)
	}
}

func main() {

	// ใส่ defer ที่ไหนก็จะไปทำงานท้ายสุด
	// กรณี ใช้ defer หลายอัน จะเกิดการทำงานแบบ LIFO หรือ Last in First Out
	fmt.Println("Welcome")
	defer loop()
	defer add(20, 12)
	add(11, 12)
	add(33, 33)
	defer fmt.Println("End")
	deferloop()
}
