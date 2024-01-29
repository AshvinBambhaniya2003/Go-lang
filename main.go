package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"


// jwttoken := 300000  connot declare this type
var jwttoken = 4500

const LoginToken string = "asjvon" 

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for pizza:")

	// comma ok or err ok

	input , _ := reader.ReadString('\n')
	fmt.Println("thank for ratindg")
	fmt.Printf("type of input %T",input)

	numRating,err  := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("added 1 to your rating",numRating+1)
	}
}