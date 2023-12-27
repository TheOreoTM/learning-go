package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	bill := newBill(name)
	fmt.Printf("Created the bill - \"%v\" \n", bill.name)

	return bill
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch strings.ToLower(option) {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip updated -", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Bill has been saved")
		fmt.Println(b.format())
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
