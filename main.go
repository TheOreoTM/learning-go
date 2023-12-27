package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string
	for _, initial := range names {
		initials = append(initials, initial[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], ""

}

func main() {
	firstName, secondName := getInitials("Usain bolt")
	fmt.Printf("Users first name is %v and second name is %v \n", firstName, secondName)

	firstName, secondName = getInitials("Mario")
	fmt.Printf("Users first name is %v and second name is %v \n", firstName, secondName)
}
