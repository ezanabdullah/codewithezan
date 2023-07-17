package main

import "fmt"

type User struct {
	name    string
	address string
	contno  int
}
type userInterface interface {
	addUser()
}

func main() {
	Ezan := User{
		name:    "abc",
		address: "bbb",
		contno:  123456,
	}
	var userinterface userInterface
	userinterface = Ezan
	userinterface.addUser()

}
func (user User) addUser() {
	fmt.Println("value is", user)
}
