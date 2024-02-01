package main

import "fmt"

func main() {
	age := 455

	if age < 18 {
		fmt.Println("you can not drive")
	} else if age > 50 {
		fmt.Println("you are very old to drive")
	} else {
		fmt.Println("you can drive")
	}

	if varnum := 1; varnum < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is grater than 10")
	}
}
