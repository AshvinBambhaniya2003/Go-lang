package main

import (
	"fmt"
	// "io"
	// "net/http"
	"net/url"
)

const Url = "https://courses.learncodeonline.in:443/learn/account/signup?signInToPay=133333&priceId=61135"

func main() {

	result, _ := url.Parse(Url)
	fmt.Println(result)


	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()

	for _, v := range qparams {
		fmt.Println(v)
	}

	partsofUrl := &url.URL {
		Scheme: "https",
		Host: "ashvinbambhaniya.netlify.com",
		Path: "/",
		RawQuery: "",
	}

	anotherurl := partsofUrl.String()
	fmt.Println(anotherurl)
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
