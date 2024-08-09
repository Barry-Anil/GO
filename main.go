package main

import (
	"fmt"
)

// GROUP A TYPES ==> strings, ints, floats, arrays, bools, and structs.
// GROUP B TYPES ==> slices, func, maps

func updateName (n *string) {
	
	// De-referencing the pointer to get the value
	*n = "Lee"

}



func main() {
	name := "Anil"

	//Pointers
	m := &name

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)


}

