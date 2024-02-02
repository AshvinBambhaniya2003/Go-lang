package main

import (
	"fmt"
	"io"
	"net/http"
)

const Url = "https://lco.dev"

func main() {

	response, err := http.Get(Url)

	checkNilerr(err)

	fmt.Printf("response type is %T",response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	checkNilerr(err)

	fmt.Println(string(databytes))

}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
