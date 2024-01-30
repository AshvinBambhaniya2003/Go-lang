package main

import "fmt"
// import "sort"
// import "os"
// import "strconv"
// import "strings"


func main()  {
	lang := make(map[string]int)
	lang["a"] = 1
	lang["b"] = 2
	lang["c"] = 3 
	lang["d"] = 4 

	fmt.Println(lang)
	fmt.Println(lang["a"])

	delete(lang,"b")
	fmt.Println(lang)

	for key,value := range lang {
		fmt.Printf("key: %v -> value %v ",key,value)
	}
	
}
