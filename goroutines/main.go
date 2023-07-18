package main

import (
	"fmt"
	"time"
)

func main() {
	go a()
	go b()
	go c()
	go d()
	time.Sleep(time.Second * 10)
}
func a() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func a is running", i)
	}
}
func b() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func b is running", i)
	}
}
func c() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func c is running", i)
	}
}
func d() {
	for i := 1; i <= 10; i++ {
		fmt.Println("func d is running", i)
	}
}
