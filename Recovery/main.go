package main

import (
	"fmt"
	"os"
)

func main() {
	dbconnection := true
	if !dbconnection {
		fatal("database not connected")
	}
	a := []string{
		"something happend",
	}
	savetodb(a)
	b := []string{}
	savetodb(b)
	c := []string{
		"happened in c",
	}
	savetodb(c)

}
func recoverpanic() {
	r := recover()
	if r != nil {

	}
}
func savetodb(record []string) {
	if len(record) < 1 {
		defer recoverpanic()
		panic("no record found")
	} else {
		fmt.Println("record =", record)
	}
}
func fatal(message interface{}) {
	fmt.Printf("message occured,message =%v", message)
	os.Exit(1)
}

// func panic(message interface{}) {
// 	fmt.Println("message occured")
// 	panic(message)
// }
