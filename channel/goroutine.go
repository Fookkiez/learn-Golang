package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Go Routine เพียงเติม go ไว้หน้าสิ่งต้องการใช้
	go f("Hello Fook")
	go f("message2")
	time.Sleep(5 * time.Second)
}
