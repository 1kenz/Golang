package interfaces

import (
	"fmt"
	"math"
)


func Demo1() {
	r := rectangle{width: 3, height: 4}
	school(r)
	
	c:= circle{radius: 5}
	school(c)
}

func school (s shape){
	fmt.Println(s.area())
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
