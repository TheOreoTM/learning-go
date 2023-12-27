package main

import "fmt"

var score = 99.5

// go run main.go greetings.go
func main() {
	sayHello("Mario")

	for _, p := range points {
		fmt.Println(p)
	}

	showScore()
}
