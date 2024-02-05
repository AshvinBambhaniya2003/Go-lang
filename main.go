// Program to pass pointer as a function argument

package main
import "fmt"

// function definition with a pointer argument
func update(num *int) {

  // dereference the pointer
  *num = 30

} 

// Program to return a pointer from a function
func display() *string {

	message := "Programiz"
  
	// returns the address of message
	return &message
  
  }

func main() {
 
  var number = 55

  // function call
  update(&number)
  
  fmt.Println("The number is", number)

  // function call
  result := display() 
  fmt.Println("Welcome to", *result)

}

