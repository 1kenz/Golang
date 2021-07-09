package slices

import "fmt"

func Demo2() {
	cities := []string{"Adelaide", "Alice Springs", "Darwin", "Melbourne", "Sydney"}
	fmt.Println(cities)
	fmt.Println(len(cities))
	fmt.Println(cap(cities))
	fmt.Println(cities[0])

	citiesCopy := make([]string, len(cities))
	fmt.Println(citiesCopy)
	copy(citiesCopy, cities)

	cities = append(cities, "Brisbane")
	fmt.Println(cities)

	fmt.Println(cities[1:4])
	fmt.Println(cities[:4])
	fmt.Println(cities[2:])

}