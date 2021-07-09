package maps

import "fmt"

func Demo1() {
	// key-valuepair
	dictionary := make(map[string]string)
	dictionary["name"] = "ken"
	dictionary["age"] = "39"
	dictionary["address"] = "Istanbul"

	fmt.Println(dictionary["name"])
	fmt.Println("Item count:", len(dictionary))
	fmt.Println("Dictionary :", dictionary)
	delete(dictionary, "address")
	fmt.Println("Item count :", len(dictionary))
	fmt.Println("Dictionary :", dictionary)

	dictionary2 := map[string]string{"name": "ken", "age": "39", "address": "Istanbul"}
	fmt.Println("Dictionary :", dictionary2)


	// loop
	for key, value := range dictionary {
		fmt.Println(key, value)
	}
	fmt.Println(dictionary)

	value, isExist := dictionary["name"]
	fmt.Println(value)
	fmt.Println("Is exists : ", isExist)

}
