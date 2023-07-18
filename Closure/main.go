package main

import "fmt"

func student() func() int {
	rollno := 0
	return func() int {
		rollno++
		return rollno
	}
}

func main() {
	StudentRollNo := student()
	fmt.Println("your roll no is", StudentRollNo())

}
