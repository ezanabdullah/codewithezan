package main

import "fmt"

func a() {
	fmt.Println("we ara in a")
}
func main() {
	go a()

}
