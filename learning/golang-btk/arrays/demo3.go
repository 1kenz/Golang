package arrays

import "fmt"

func Demo3() { numbers := [10]int{10, 22, 33, 64, 35, 76, 17, 88, 91, 10}; fmt.Println(numbers)

bigger := 0
for i := 0; i < len(numbers); i++ {
	
	if numbers[i] > bigger {
		bigger = numbers[i]
	}
}
fmt.Println("Bigger number is",bigger)
}