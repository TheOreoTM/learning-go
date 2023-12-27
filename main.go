package main

import (
	"fmt"
)

func main() {
	// fmt.Println() prints and adds a \n at the end fmt.Print() doesnt

	const text string = "hello there"
	const age int8 = 20
	const name string = "Foo"

	fmt.Println("my age is", age, "and my name is", name)

	// Printf - (formatted strings) %_ = format specifier

	fmt.Printf("my age is %q and my name is %q \n", age, name)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}
