package main

import "fmt"

//Struct
type Student struct {
	Name           string
	Class          int
	RollNumber     int
	StudentAddress Address
}

//Nested-Struct
type Address struct {
	Lane1   string
	Lane2   string
	Post    int
	Dist    string
	Village string
}
type Phone struct {
	Name      string
	Brand     string
	IMEI      int
	Configure Config
}
type Laptop struct {
	BasicInfo
	Name          string
	Brand         string
	Model         int
	Configuration Config
}
type Config struct {
	Ram       int
	Rom       int
	Processor int
}
type BasicInfo struct {
	Brand string
	Model string
}

func main() {
	//Non-Primitive approach
	var student Student
	student.Class = 10
	student.Name = "vishal"
	student.RollNumber = 11
	student.StudentAddress.Lane1 = "Bijnor"
	student.StudentAddress.Lane2 = "Qazipada"
	student.StudentAddress.Dist = "bol"
	student1 := Student{Class: 12, Name: "eeeee", RollNumber: 125, StudentAddress: Address{
		Lane1:   "abc",
		Lane2:   "def",
		Post:    246701,
		Dist:    "ssss",
		Village: "dddd",
	}}
	var Ezan Phone
	Ezan.Brand = "vivo"
	Ezan.IMEI = 1212121
	Ezan.Name = "v19"
	Ezan.Configure.Processor = 12
	Ezan.Configure.Ram = 12
	Ezan.Configure.Rom = 128
	Ezne := Phone{Brand: "vivo", Configure: Config{
		Ram: 10,
		Rom: 12,
	}}
	var Ezan1 Laptop
	Ezan1.BasicInfo.Brand = "lenovo"
	Ezan1.BasicInfo.Model = "ThinkPad"
	Ezan1.Name = "aaa"
	Ezan1.Model = 123

	fmt.Println("hello world", student, student1, Ezan, Ezne, Ezan1)
	//Interface
	val1 := 123
	val2 := "anv"
	//Interface
	var interfaceExample interface{}
	interfaceExample = val1
	fmt.Println("InterfaceExample", interfaceExample)
	interfaceExample = val2
	fmt.Println("InterfaceExample", interfaceExample)
	interfaceExample = false
	fmt.Println("InterfaceExample", interfaceExample)
}
