package main

import "fmt"

func main() {
	fmt.Println("hello from main")
	greet()

	// func greest()  {
	// 	fmt.Println("namste")
	// }  -->this is not allow

	result := sum(4,5)

	fmt.Println(result)

	// proresult := prosum(4,5,5,1,0,5)

	proRes, proMessage := prosum(4,5,6,1,7)

	fmt.Println(proRes,proMessage)
}

func sum(a int , b int) int {
	return a + b
}

func prosum(values ...int) (int, string){
	total := 0

	for _, val := range values{
		total = total + val
	}

	return total,"hi from pro result function"
}

func greet()  {
	fmt.Println("namste")
}