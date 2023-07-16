package main

import "fmt"

func main() {
	//declaring of map
	m := make(map[string]int)
	//save data in map
	m["abc"] = 12
	m["pqr"] = 13
	//fetch data from map
	val, ok := m["ijk"]
	fmt.Println("val is", val, "ok is", ok)
	if ok {
		fmt.Println("value is found", val)
	} else {
		fmt.Println("value is not found")
	}
}
