package stringfunctions

import (
	"fmt"
	"strings"
)


func Demo2() { 
	name := "John"; 
	fmt.Println(strings.HasPrefix(name, "Jo"))
	fmt.Println(strings.HasSuffix(name, "Jo"))
	fmt.Println(strings.Index(name, "Jo"))
	
	letters := []string{"k","e","n","a","n"}
	result := strings.Join(letters, "-")
	fmt.Println(result)

	fmt.Println(strings.Replace(result, "-", "*",-1))

	fmt.Println(strings.Split(result,"-"))
	fmt.Println(strings.Repeat(result, 3))

}