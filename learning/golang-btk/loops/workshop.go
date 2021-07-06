package loops

import (
	"fmt"
)


func Workshop() {

	number := 55
	
	var prediction int = 56


	if number < prediction {
		fmt.Println("Lower")
	} else if number > prediction {
		fmt.Println("Higher")
	} else {
		fmt.Println("Thats correct!!!")
	}



	// fmt.Println(prediction)
	// fmt.Println(number)
}