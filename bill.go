package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.00,
	}

	return b
}

// format the bill
func (b *bill) format() string {
	formattedString := "Bill breakdown: \n\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		formattedString += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	formattedString += "\n"

	// add tip
	formattedString += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	// total
	formattedString += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total)

	return formattedString
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
