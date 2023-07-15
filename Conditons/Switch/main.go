package main

import "fmt"

func main() {
	daysofweek := 11
	switch daysofweek {
	case 1:
		{
			fmt.Println("monday")
		}
	case 2:
		{
			fmt.Println("tuesday")
		}
	case 3:
		{
			fmt.Println("wednesday")
		}
	default:
		{
			fmt.Println("not a weekday")
		}
	}
}
