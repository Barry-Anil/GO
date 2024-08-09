package main

import (
	"fmt"
)


func main() {
	bill := newBill("arry's bill")
	fmt.Println(bill.format())
}

