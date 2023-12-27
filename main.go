package main

import "fmt"

// return "piper"
func updateName(name string) string {
	name = "piper"
	return name
}

// adds ice cream to the menu
func updateMenu(menu map[string]float64) {
	menu["ice cream"] = 12.99
}

func main() {
	// group A types => strings, ints, bools, floats, arrays, structs (non pointer values)
	name := "carl"

	updateName(name)

	fmt.Println(name) // carl

	name = updateName(name)

	fmt.Println(name) // piper

	// group B types => slices, maps, functions (pointer wrapper values)
	menu := map[string]float64{
		"pie":    5.49,
		"coffee": 8.99,
	}

	updateMenu(menu)

	fmt.Println(menu) // map[coffee:8.99 ice cream:12.99 pie:5.49]
}
