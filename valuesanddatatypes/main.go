package main

import "fmt"

func main() {
	//12 -int
	//12.2222-float
	//true/false-boolean
	//c-char/string
	//coding concepts-string
	var a int
	a = 10
	var b float32
	b = 25.456
	var flags bool
	flags = true
	var name string
	name = "ezaneeee"
	fmt.Printf("a=%d,b=%f,flags=%v,name=%s", a, b, flags, name)
	//const number1 =12345
	//number1 =25843
	//Pointer In Golang
	var addressofnumber *int
	addressofnumber = &a
	fmt.Printf("addressofnumber=%d", addressofnumber)
	var decinumb float64
	decinumb = 12.56544
	var addressoffloat *float64
	addressoffloat = &decinumb
	fmt.Printf("addressoffloat=%v", addressoffloat)
	var addressofflag *bool
	addressofflag = &flags
	fmt.Printf("addressofflag=%v", addressofflag)
	var addressofstring *string
	addressofstring = &name
	fmt.Printf("addressofstring=%s", *addressofstring)
	//types of declaration
	number12 := 123
	fmt.Println(number12)

}
