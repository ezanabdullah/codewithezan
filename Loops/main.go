package main

import "fmt"

func main() {
	a := 2
	for i := 0; i <= 10; i++ {
		b := a * i
		fmt.Println(b)
	}
	a = 10
	condition := true
	for condition {
		a++
		if a == 101 {
			condition = false
		} else {
			fmt.Println(a)
		}
	}
	//range in golang
	ab := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(ab); i++ {
		fmt.Println(ab[i])
	}
	for i, j := range ab {
		fmt.Println("at i=", i, "value is", j)
	}
	//map declaration
	m := make(map[string]int)
	m["abc"] = 2
	m["def"] = 3
	for key, val := range m {
		fmt.Println("key is", key, "val is", val)
	}
}
