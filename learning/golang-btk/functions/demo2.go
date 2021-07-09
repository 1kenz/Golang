package functions

func Calculate(number1, number2 int) (int, int, int, float32) {
	summary := number1 + number2
	difference := number1 - number2
	multiply := number1 * number2
	divide := float32(number1 / number2)
	return summary, difference, multiply, divide
}