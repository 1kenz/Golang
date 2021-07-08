package loops

import "fmt"

// friend numbers


func Workshop3() {
	number1 := 11  // 220
	number2 := 44  // 284
	sum1 := 0
	sum2 := 0
	
	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			sum1 = sum1 + i
		}

		if number2%i == 0 {
			sum2 = sum2 + i
		}
	} 
	
		if sum1 == number2 && sum2 == number1 {
			fmt.Println("This numbers are friend")
		} else {
			fmt.Println("These numbers are not friend")
		}
}