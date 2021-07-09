package for_range

import "fmt"

func Demo1() {
	cities := []string{"İstanbul", "Ankara", "İzmir", "Muğla", "Sydney"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for i, city:= range cities {
		fmt.Println(i, city)
	}
	
	for _, city:= range cities {
		fmt.Println(city)
	}
}