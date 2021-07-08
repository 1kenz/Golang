package loops

// 1. User input a number
// 2. For number in range [0, 100)
// 3. If prime number is only divide by itself and 1
// 4. Print number
// 5. Exit

import "fmt"

func Workshop2() {
	var number int
	fmt.Println("Please input a number")
	// fmt.Scanf("%d", &number)
	fmt.Scanln(&number)

	isPrime := true
	for i := 2; i < number; i++ {
		
		if number%i == 0 {
			isPrime = false
		}
	}
	if isPrime {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "is not a prime number")
	}
}