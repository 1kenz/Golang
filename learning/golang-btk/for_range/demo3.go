package for_range

import (
	"fmt"
)

func Demo3() {
	dictionary:= map[string]string{"book": "sözlük", "table": "masa"}
	for key, value := range dictionary {
		fmt.Println(key)
		fmt.Println(value)
}
}