package conditionals

import "fmt"

func Workshop() {
	// number1 := 555
	// number2 := 55
	// number3 := 22

	number1, number2, number3 := 555, 777, 10

	max := number1

	if number2 > max {
		max = number2
	} else if number3 > max {
		max = number3
	} 

	fmt.Printf("Bigger number : %v", max)
}