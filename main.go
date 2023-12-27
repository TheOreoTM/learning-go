package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":   4.99,
		"pie":    7.99,
		"salad":  6.99,
		"toffee": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type

	phonebook := map[int]string{
		1208021: "mario",
		3209823: "luigi",
		2398102: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2398102])

	phonebook[3209823] = "bowser"

	fmt.Println(phonebook[3209823])

}
