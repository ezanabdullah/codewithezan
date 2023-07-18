package main

import "fmt"

type User struct {
	name    string
	address string
	contno  int
}
type Student struct {
	name   string
	rollno int
	class  int
}

func main() {
	var ezan User
	ezan.name = "abc"
	ezan.address = "xyz"
	ezan.contno = 123456
	ezan.addUser(20)
	ezan1 := Student{
		name:   "qwerty",
		rollno: 1234,
		class:  12,
	}
	ezan1.addStudent(15)
	ezan1.addStudent1(16)
	ezan2 := User{
		name:    "qqqqqq",
		address: "sssssss",
		contno:  789456123,
	}
	ezan2.addUser(30)
}

// Method function
func (user *User) addUser(a int) {
	user.name = "asdcvb"
	fmt.Println("The values are", user)
}
func (student Student) addStudent(a int) {
	student.name = "1111111"
	fmt.Println("The values are", student)
}
func (student Student) addStudent1(a int) {
	fmt.Println("the values are", student)
}
