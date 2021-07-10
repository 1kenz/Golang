package deferstatement

import "fmt"

func Demo2(number int32) string {
	defer DeferFunc()
	
	if number%2 == 0 {
		return "even"
	} else if number%2 != 0 {
		return "odd"
	}
	return "error"
}

func Test() {
	result := Demo2(9)

	fmt.Println(result)
}

func DeferFunc() {
	fmt.Println("Finished!!!")
}
