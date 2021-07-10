package errorhandling

import (
	"fmt"
	"os"
)


func Demo1() {
	f,err := os.Open("demo11.txt")  // return file, error

	// type assertion
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
		fmt.Println("File not found!", pErr.Path) // handle error
		return
		} else {
			fmt.Println("File not found!")
			return
			} 
	} else {
		fmt.Println(f.Name()+ " opened!")
		return
	}	
}