package structs

import "fmt"

type customer struct {
	firstName string
	lastName string
	age      int
}

func (c customer) save() {
	fmt.Println(c.firstName, "Added.")
}

func (c customer) update() {
	fmt.Println(c.firstName, "updated.")
}

func Demo2() {
	c := customer{
		firstName: "John",
		lastName: "Nash",
		age:      30,
	}
	c.save()
	c.update()
}