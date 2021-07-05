package main

import "fmt"

// func main() {

// 	//! Print
// 	// fmt.Println("Hello, world!!!")

// 	//! Variables
// 	// var x int = 5
// 	// var y int = 7
// 	// var sum int = x + y

// 	//! Shorthand
// 	// x := 5
// 	// y := 7
// 	// sum := x + y

// 	// fmt.Println(sum)

// 	//! If Statement
// 	// if x > 6 {
// 	// 	fmt.Println("More than 6 :)")
// 	// } else if x < 6 {
// 	// 	fmt.Println("Less than 6 ")
// 	// } else{
// 	// 	fmt.Println("???")
// 	// }

// 	//! Arrays
// 	// var a [5]int  // define all variables 0 thats mean number or string same result but type is different
// 	// a[2] = 7

// 	// array2 := [5]int{1,2,3,4,5}
// 	// array3 := []int{1,2,3,4,5}
// 	// array3 = append(array3, 13)  // return new array

// 	// var s [5]string

// 	// fmt.Println(a)
// 	// fmt.Println(s)
// 	// fmt.Println(array2)
// 	// fmt.Println(array3)

// 	//! Maps -- key-value  -- like python dictionary
// 	// vertices := make(map[string]int)  // python dictionary key-value

// 	// vertices["triangle"] = 2
// 	// vertices["square"] = 3
// 	// vertices["dodecagon"] = 12

// 	// delete(vertices, "triangle")

// 	// fmt.Println(vertices)
// 	// fmt.Println(vertices["triangle"])

// 	//! Loops

// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	//! While Loop
// 	// i := 0

// 	// for i < 5 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }

// 	//! Range
// 	// array4 := []string{"a", "b", "c"}

// 	// for index, value := range array4 {
// 	// 	fmt.Println("index", index, "value", value)
// 	// }

// 	// for index, value := range array4 {
// 	// 	fmt.Println("index", index, "value", value) }

// 	// //! Map
// 	// array5 := make(map[string]string)
// 	// array5["name"] = "ken"
// 	// array5["surname"] = "den"

// 	// for key, value := range array5 {
// 	// 	fmt.Println("key", key, "value", value)
// 	// }

// 	//! Call a function
// 	// result := sum(2,3)
// 	// fmt.Println(result)

// 	//! Call sqrt function -- multi value return
// 	// result, err := sqrt(9)
// 	// // result, err := sqrt(-1)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	fmt.Println(result)
// 	}

// }

//! Add a new function
// func sum(x int, y int) int {
// 	return x + y
// }

//! Multi return
// func sqrt(x float64) (float64, error) {

// 	if x < 0 {

// 		return 0, errors.New("Undefined for negative numbers")
// 	}

// 	return math.Sqrt(x), nil
// }

//! Struct type
// type person struct {
// 	name string
// 	age int
// }

// func main(){
// 	p := person{name:"Ken", age:39}
// 	fmt.Println(p)
// 	fmt.Println(p.age)
// 	fmt.Println(p.name)
// }

//! Pointers
func main(){
	i := 7
	fmt.Println(i)
	fmt.Println(&i)  // if use & define i reference address we use update variable etc. do not change reference address
	inc(&i)
	fmt.Println(i)
	fmt.Println(&i)
}

func inc(x *int){  // not create new reference address if use *
	*x++
}