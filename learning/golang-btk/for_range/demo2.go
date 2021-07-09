package for_range

import "fmt"

// 1-10 sum for-range only odd numbers


func Demo2() {
	numbers:=[]int{1,2,3,4,5,6,7,8,9,10}
	sum := 0
	for _, number := range numbers {
		if number%2 != 0 {
			sum += number
		}
	}
	fmt.Println("Summary",sum)
}