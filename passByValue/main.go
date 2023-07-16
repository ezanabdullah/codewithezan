package main

import "fmt"

func main() {
	a := 5
	b := 6
	fmt.Println("a/b is", a/b)
	fmt.Println("add ans is", add(a, b))
	fmt.Println("sub ans is", sub(10, 5))
	fmt.Println("multiply answer is ", mul(10, 10))
	divide, err := div(10, 15)
	if err != nil {
		fmt.Println("err is", err)
	} else {
		fmt.Println("division is", divide)
	}

}

//for ading two nos
func add(a, b int) int {
	c := a + b
	return c
}

//for subtracting
func sub(a, b int) int {
	return a - b
}

//for multiply
func mul(a, b int) int {
	return a * b
}

//for div
func div(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	} else {
		return 0, fmt.Errorf("dividing with zero is not possible")
	}
}
