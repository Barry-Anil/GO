package main

import (
	"fmt"
)

// GROUP A TYPES ==> strings, ints, floats, arrays, bools, and structs.
// GROUP B TYPES ==> slices, func, maps

// func updateName (n string) {
// 	n = "Lee"
// }

// func main() {
// 	name := "Anil"

// 	updateName(name)
// 	// this will print anil as functions stores only the copy of the variable name.
// 	fmt.Println(name)

// }

func updateName (n string) string {
	n = "Lee"
	return n
}

func updateMenu (x map[string]float64) {
	x["coffee"] = 12.99
}

func main() {
	name := "Anil"

	name = updateName(name)
	// this will print the update value for the var name as we are returning the value from the function.
	fmt.Println(name)

	menu := map[string]float64{
		"pie" : 6.99,
		"choco" : 4.99,
	}

	updateMenu(menu)
	fmt.Println(menu)


}

