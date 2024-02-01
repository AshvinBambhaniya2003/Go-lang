package main

import "fmt"

func main() {
	ashvin := User{"ashvin", "ashvin@gmai.com", true, 20}
	fmt.Println(ashvin)
	fmt.Printf("details of ashvin is %+v\n", ashvin)
	fmt.Printf("name is %v and email is %v", ashvin.Name, ashvin.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}