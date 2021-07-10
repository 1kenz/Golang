package channels

func EvenNumbers(EvenNumbersCn chan int) {
	total := 0
	for i := 0; i < 10; i += 2 {
		total += i
		// fmt.Println("Even numbers :", i)
		// fmt.Println("Total numbers :", total)
		// time.Sleep(1 * time.Second)
	}
	EvenNumbersCn <- total
}

func OddNumbers(OddNumbersCh chan int) {
	total := 0
	for i := 1; i < 10; i += 2 {
		total += i
		// fmt.Println("Odd numbers :", i)
		// time.Sleep(1 * time.Second)
	}
	OddNumbersCh <- total
}