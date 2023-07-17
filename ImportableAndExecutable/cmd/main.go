package main

import (
	"fmt"

	"github.com/ezanabdullah/codewithezan/ImportableAndExecutable/calculator"
)

func main() {
	usr := calculator.User{}
	usr.Name = "shivam"

	fmt.Println(usr)
	fmt.Println("the sum is", calculator.Add(10, 15))
	fmt.Println("the sub is ", calculator.Sub(15, 10))
	fmt.Println("the mul is", calculator.Mul(2, 4))
	fmt.Println("the div is", calculator.Div(15, 5))

}
