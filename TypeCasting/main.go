package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	a = 16
	var b int16
	b = int16(a)
	fmt.Println("the value is", b)
	//Integer To String Converter
	var c string
	c = fmt.Sprintf("%d", b)
	fmt.Println("the value is ", c)
	//String To Integer Converter
	num := "123456789"
	num1, _ := strconv.Atoi(num)
	fmt.Println("the value is ", num1)

}
