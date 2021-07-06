package conditionals

import "fmt"

func Demo2() {
	var account float32 = 1000
	var draw float32 = 900

	if account < draw {
		fmt.Println("No sufficient money in your account")
	} else if account == draw {
		fmt.Println("Preparing your money...")
		account = account - draw
	} else {
		fmt.Println("Preparing your money...")
		account = account - draw
	}
	fmt.Println("Your money is ready. Now your balance : " + fmt.Sprintf("%v", account))
	fmt.Printf("Your money is ready. Now your balance : %v", account)
}