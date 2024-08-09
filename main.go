package main

import (
	"fmt"
)


func main() {
	bill := newBill("arry's bill")
	bill.addItem("onion soup", 4.50)
	bill.addItem("Momos", 4.50)
	bill.addItem("Noodles", 4.50)
	bill.addItem("Hot soup", 4.50)

	bill.updateTip(10)
	fmt.Println(bill.format())
}

