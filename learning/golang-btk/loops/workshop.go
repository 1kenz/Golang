package loops

import (
	"fmt"
)


func Workshop() {

	number := 55
	
	prediction := 0

	fmt.Println("Start prediction!")
	fmt.Scanln(&prediction)


	for number != prediction {

		if number < prediction {
			fmt.Println("Lower")
			fmt.Scanln(&prediction)
		} else if number > prediction {
			fmt.Println("Higher")
			fmt.Scanln(&prediction)
		}
	}
	fmt.Println("Thats correct!!!")

				// fmt.Println(prediction)
				// fmt.Println(number)
}