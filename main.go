package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a New Bill: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an Option (a - Add Item, s - Save the Bill, t - Add Tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter the item name: ", reader)
		price, _ := getInput("Enter the Item Price($): ", reader)
		fmt.Println(name, price)
	case "s":
		tip, _ := getInput("Enter the Tip Amount($): ", reader)
		fmt.Println(tip)
	case "t":
		fmt.Println("You Choose t")
	default:
		fmt.Println("You Choose an Inalid Option...")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill.format())
}
