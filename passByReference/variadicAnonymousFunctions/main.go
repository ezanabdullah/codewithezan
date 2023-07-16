package main

import "fmt"

func main() {
	fmt.Println("the sum is", add(4, 5, 6))
	//anonymous functions
	func(a int) int {
		fmt.Println("anonymous func is running")
		return 0
	}(55)
}

//1st way variadic function
func add(b ...int) int {
	sum := 0
	for _, val := range b {
		sum = sum + val
	}
	return sum

}

//Anonymous Function
