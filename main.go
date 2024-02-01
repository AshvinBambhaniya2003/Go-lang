package main

import "fmt"

func main() {
	ashvin := User{"ashvin", "ashvin@gmai.com", true, 20}
	fmt.Println(ashvin)
	// fmt.Printf("details of ashvin is %+v\n", ashvin)
	// fmt.Printf("name is %v and email is %v", ashvin.Name, ashvin.Email)

	ashvin.GetStatus()
	ashvin.newMail()
	fmt.Println(ashvin)  //original value is change only the copy of object is pass
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("is user active: ",u.Status)
}

func (u User) newMail() {
	u.Email = "ashvinbambhaniya@gmail.com"
	fmt.Println("email is :",u.Email)
}