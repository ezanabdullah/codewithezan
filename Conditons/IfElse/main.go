package main

import "fmt"

func main() {
	num := 1
	if num >= 10 && num >= 8 {
		fmt.Println("hiiii")
	} else if num == 0 || num <= 10 {
		fmt.Println("bye")
	}
}
