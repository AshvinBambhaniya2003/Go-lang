package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "this is need to go in file ans my also "

	file, err := os.Create("./mylcogifile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkNilerr(err)

	length, err := io.WriteString(file, content)

	checkNilerr(err)

	fmt.Println(length)
	defer file.Close()
	readFile("./mylcodgifile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilerr(err)

	fmt.Println("data byte is ", databyte)
	fmt.Println("data is:", string(databyte))
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
