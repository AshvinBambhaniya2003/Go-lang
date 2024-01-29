package main

import "fmt"

// jwttoken := 300000  connot declare this type
var jwttoken = 4500

const LoginToken string = "asjvon" 

func main()  {
	var username string = "ashvin"
	fmt.Println(username)
	fmt.Printf("variable of type %T \n",username)

	var isLogin bool = false
	fmt.Println(isLogin)
	fmt.Printf("variable of type %T \n",isLogin)

	// defult value of int
	var anotherint int
	fmt.Println(anotherint)
	fmt.Printf("variable of type %T \n",anotherint)

	numberoduser := 4000
	fmt.Println(numberoduser)

	fmt.Println(jwttoken)
	fmt.Println(LoginToken)

}