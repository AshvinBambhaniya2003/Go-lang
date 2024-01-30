package main

import "fmt"
import "sort"
// import "os"
// import "strconv"
// import "strings"


func main()  {
	var fruits = []string{}
	fmt.Printf("%T \n",fruits)

	fruits = append(fruits, "a", "b", "c" , "d", "e")
	fmt.Println(fruits)

	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	highScore := make([]int,4)
	highScore[0] = 5
	highScore[1] = 2
	highScore[2] = 1
	highScore[3] = 4
	// highScore[4] = 1  --> this will show error

	highScore = append(highScore, 3,9,7)

	fmt.Println(highScore)

	sort.Ints(highScore)  //sort integer slice

	fmt.Println(highScore)

	var corses = []string{"a","b","c","d","e","f"}
	fmt.Println(corses)

	index := 2

	corses = append(corses[:index],corses[index+1:]...)
	fmt.Println(corses)

}
