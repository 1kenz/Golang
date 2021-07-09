package arrays

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "Istanbul"
	cities[2] = "Izmir"
	cities[3] = "Bursa"
	cities[4] = "Adana"
	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i])
	}
}
