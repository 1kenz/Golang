package for_range

import (
	"fmt"
)

func Demo3() {
	dictionary:= map[string]string{"book": "sözlük", "table": "masa"}
	for key, value := range dictionary {
		fmt.Print(key)
		fmt.Print(" : ")
		fmt.Println(value)
}
}