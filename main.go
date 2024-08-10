package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The Price must be a number.")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("Item Added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter the Tip Amount($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The Tip must be a number.")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip Added - ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("You Choose to save the bill", b)

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
