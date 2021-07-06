package conditionals

import "fmt"

func Demo1() {
	var account float32 = 1000
	var draw float32 = 900

	if account > draw {
		fmt.Println("No sufficient money in your account")
	}
}