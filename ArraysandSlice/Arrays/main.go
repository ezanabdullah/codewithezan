package main

import "fmt"

func main() {
	var name [10]string
	name[0] = "aaa"
	name[1] = "bbb"
	name[2] = "ccc"
	name[3] = "ddd"
	name[4] = "e"
	name[5] = "f"
	name[6] = "g"
	name[7] = "h"
	name[8] = "i"
	name[9] = "j"
	fmt.Println("name is", name)
	a := [3]string{"abc", "def", "ghi"}
	fmt.Println(a)
	//slice
	var students []string
	students = append(students, "ayush")
	students = append(students, "chunty")
	fmt.Println("students is", students)
	b := []string{"apple", "ball", "cat"}
	fmt.Println(b)

}
