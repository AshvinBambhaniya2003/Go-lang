package main
import "fmt"


// create struct
type Weather struct{
	city string
	temperature int
}   

// declare a struct Person	  
type Person struct {
	name string
	age int
}

func main() {

  // instance of the struct Person
  person1 := Person{"John", 25}

  // create a struct type pointer that
  // stores the address of person1
  var ptr *Person
  ptr = &person1

  // print struct instance
  fmt.Println(person1)

  // print the struct type pointer
  fmt.Println(ptr)
  fmt.Println(*ptr)
  fmt.Println((*ptr).name)
  fmt.Println(ptr.name)


  // create struct variable
  weather := Weather{"California", 20}
  fmt.Println("Initial Weather:", weather)

  // create struct type pointer
  ptr1 := &weather

  // change value of temperature to 25
  ptr1.temperature = 25 
//   (*ptr1).temperature = 25   //also work 

  fmt.Println("Updated Weather:", weather)


}


