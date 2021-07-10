package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save(){ 
	fmt.Println(p.productName, "saved to db", )
	defer Log()
}

func Log(){
	fmt.Println("logged!")
}

func Demo3(){
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	fmt.Println("Success")

	p = product{productName: "Mouse", unitPrice: 25}
	defer p.save()
}
