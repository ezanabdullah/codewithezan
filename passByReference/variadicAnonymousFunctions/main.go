package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30
	fmt.Println("the sum is", add(&a, &b, &c))

}
func add(b ...*int) int {
	var sum *int
	var val *int
	for _, val = range b {
		*sum = *sum + *val
	}
	return *sum
}
