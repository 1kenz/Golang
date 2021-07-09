package functions

func SumVariadic(numbers ...int) int {
	summarize := 0
	// for i:=0,; i<len(numbers); i++ {
	// 	summarize += number[i]
	// }
	// return summarize


	for _, number := range numbers {
		summarize += number
	}
	return summarize
}