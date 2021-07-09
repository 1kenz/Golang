package slices

import "fmt"

func Demo1() {
	names := make([]string, 5)
	names[0] = "Ken"
	names[1] = "Den"
	names[2] = "Rose"
	names[3] = "Ker"
	names[4] = "Met"
	names = append(names, "Neva")
	fmt.Println(names)
	fmt.Println(len(names))
}