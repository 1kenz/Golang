package interfaces

import "fmt"

type Mortgage struct {  // first credit
	Amount  float64
	address string
	rate    float64
	loan    int
}

type Car struct {  // second credit
	Amount  float64
	carInfo string
	rate    float64
	loan    int
}

type CreditCalculator interface {  // interface -- common methods
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {  // implementation of interface
	return m.Amount * (m.rate / 100) / float64(m.loan)
}

func (c Car) Calculate() float64 {  // implement interface 
	return c.Amount * (c.rate / 100) / float64(c.loan)
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {  // polymorphic
	paymentTotal := 0.0

	for _, credit := range credits {
		paymentTotal += credit.Calculate()
	}

	// for i := 0; i < len(credits); i++ {
	// 	paymentTotal = paymentTotal + credits[i].Calculate()
	// }

	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, Amount: 100000, loan: 12}
	credit2 := Car{rate: 10, Amount: 50000, loan: 12}
	credit3 := Car{rate: 15, Amount: 75000, loan: 24}

	credits := []CreditCalculator{credit1, credit2, credit3}  // polymporphic
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Total payment:", total)
}
