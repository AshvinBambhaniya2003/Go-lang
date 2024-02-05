package main
import "fmt"

// empty interface as function parameter
func displayValue(i interface {}) {
  fmt.Println(i)
}

// function with an empty interface as its parameter
func displayValueMultiple(i... interface {}) {
	fmt.Println(i)
  }

func main() {

  a := "Welcome to Programiz"
  b := 20
  c := false

  // pass string to the function 
  displayValue(a)

  // pass integer number to the function
  displayValue(b)

  // pass boolean value to the function
  displayValue(c)

  // function call with a single parameter
  displayValueMultiple(a)

  // function call with 2 parameters
  displayValueMultiple(a, b)

  // function call with 3 parameters
  displayValueMultiple(a, b, c)

}