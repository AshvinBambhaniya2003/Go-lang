package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)

	func() {
		fmt.Println("hello world first class function")
	}()  // here () is instantly call the anonymous functions 

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
	
}
