package main

import "fmt"

type Age interface {
	int64 | int32 | float32 | float64 
}

func newGenericFunc[age int64 | float64](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}


func newGenericFuncAny[age any](myAge age) {
	// val := int(myAge) + 1  // type casting : this is also not allowed 
	// fmt.Println(val)
	fmt.Println(myAge)
}

func newGenericFuncInterface[age Age](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
	// newGenericFunc(testString)  --> this is not allowed

	var testString string = "Elliot"

	newGenericFuncAny(testAge)
	newGenericFuncAny(testAge2)
	newGenericFuncAny(testString)
}