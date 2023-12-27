package main

import "fmt"

func updateName(name *string) {
	*name = "piper"
}

func main() {
	name := "carl"

	// updateName(name)

	// fmt.Println("memory address of name is:", &name)

	m := &name

	fmt.Println("value at memory address", *m)

	updateName(m)

	fmt.Println("value at memory address", *m)
}
