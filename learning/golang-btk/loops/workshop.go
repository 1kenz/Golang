package loops

import (
	"fmt"
)


func Workshop() {

	number := 55
	
	prediction := 0

	predictionCounter := 0

	fmt.Println("Start prediction!")
	fmt.Scanln(&prediction)
	predictionCounter += 1


	for number != prediction {

		if number < prediction {
			fmt.Println("Lower")
			fmt.Scanln(&prediction)
			predictionCounter += 1

		} else if number > prediction {
			fmt.Println("Higher")
			fmt.Scanln(&prediction)
			predictionCounter += 1

		}
	}

	predictionState :=""

	if predictionCounter > 0 && predictionCounter <= 3 {
		predictionState = "perfect"
	} else if predictionCounter > 3 && predictionCounter <= 6 {
		predictionState = "good"
	} else {
		predictionState = "bad"
	}

	fmt.Printf("Thats correct!!! %v prediction : %v", predictionCounter, predictionState)

				// fmt.Println(prediction)
				// fmt.Println(number)
}