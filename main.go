package main

import "fmt"
// import "bufio"
// import "os"
// import "strconv"
// import "strings"


func main()  {
	var demoptr *int 
	fmt.Println(demoptr)

	mynumer  := 45
	var ptr = &mynumer
	fmt.Println(ptr,*ptr)

	*ptr = *ptr + 45
	fmt.Println(mynumer)
}