package main

import (
	"fmt"
	"os"
)

func main() {
	pass := "123456"
	if len(pass) <= 8 {
		err := fmt.Errorf("pass length is too short password=%v", pass)
		panic(err)
	} else {
		password(pass)
	}
	//fatal Error
	dbconnection := true
	if !dbconnection {
		fatalerror("cannnot connect with db")
	}
}
func password(pass string) {
	fmt.Println("match the criteria")
}

func fatalerror(message string) {
	fmt.Printf("some error occurs which is causing fatal error ,message=%v\n", message)
	os.Exit(1)

}
