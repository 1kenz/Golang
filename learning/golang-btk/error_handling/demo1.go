package errorhandling

import (
	"fmt"
	"os"
)


func Demo1() {
	 f,err := os.Open("demo1.txt")  // return file, error

	 if err != nil {
		fmt.Println("File not found!") // handle error
	} else {
		fmt.Println(f.Name()+ " opened!")
	}	
}
