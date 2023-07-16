package main

import "fmt"

func main() {
	c := 5
	d := 5
	addbyref(&c, d)
	fmt.Println("value of c is", c, "value of d is", d)
}
func addbyref(a *int, b int) {
	*a = *a + 1
	b = b + 1
}
