package stringfunctions

import (
	"fmt"
	"strings"
)


func Demo1() {
	name := "John"
	fmt.Println(strings.Count(name, "o"))
	fmt.Println(strings.Count(name, "a"))
	fmt.Println(strings.Count(name, "j"))  // 0 because of golang is case sensitive
	fmt.Println(strings.Contains(name, "h"))  
	fmt.Println(strings.Index(name, "n"))  
	fmt.Println(strings.ToLower(name))  
	fmt.Println(strings.ToUpper(name))  
}