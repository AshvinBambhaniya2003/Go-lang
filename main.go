package main

import "fmt"

func main() {
	days := []string{"sun", "mon", "tue", "wen", "thur", "friday", "saturday"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	roughval := 1

	for roughval < 10 {

		if roughval == 5 {
			roughval++
			// continue
			// break
			goto lco
		}
		fmt.Println(roughval)
		roughval++
	}
	fmt.Println("outer of lco")
lco:
	fmt.Println("in goto statement")

	fmt.Println("outer of lco")

}