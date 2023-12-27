package main

import "fmt"

func main() {
	myBill := newBill("Marios bill")

	myBill.updateTip(10)
	myBill.addItem("hot tub", 300)

	fmt.Println(myBill.format())
}
