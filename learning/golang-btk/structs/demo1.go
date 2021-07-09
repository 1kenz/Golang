package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 500, "Dell"})
	fmt.Println(product{Name:"Bilgisayar", Price:1000})
}

type product struct {
	Name  string
	Price float32
	Brand string

}
