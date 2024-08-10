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

func promptOptions (b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an Option (a - Add Item, s - Save the Bill, t - Add Tip): ", reader)

	fmt.Println(opt)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill.format())
}

