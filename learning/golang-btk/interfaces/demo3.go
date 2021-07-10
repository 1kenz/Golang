package interfaces

import "fmt"

func isValid(i interface{}) {
	number, ok := i.(int)  // ok => true , value: 7 ,  
						   // ok=> false , value: 0 -- default int value
	fmt.Println(number, ok)
}

func Demo3() {
	var number1 interface{} = 7
	isValid(number1)

	var number2 interface{} = "ken"
	isValid(number2)
}
