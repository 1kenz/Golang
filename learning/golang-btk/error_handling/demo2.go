package errorhandling

import (
	"errors"
	"fmt"
)

func prediction(predict int) (string,error) {

	number := 77

	if predict < 1 || predict > 100 {
		return "", errors.New("Prediction value out of range (1-100)")
	}

	if predict == number {
	return "Correct", nil
	} else if predict < number {
	return "Low", nil
	} else {
	return "High", nil
	}
}

func Demo2(){
	message, err := prediction(77)
	fmt.Println(message)
	fmt.Println(err)
	fmt.Println(prediction(33))
	fmt.Println(prediction(50))
	fmt.Println(prediction(88))
	fmt.Println(prediction(111))
}