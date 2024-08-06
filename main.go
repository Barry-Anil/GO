package main

import "fmt"

func main() {
	fmt.Println("Hello, Ninjas!!")
	var name string = "Barry"
	var nameTwo = "lee"
	var nameThree string

	fmt.Println(name, nameTwo, nameThree)

	name = "anil"
	nameThree = "Barry Anil"

	fmt.Println(name, nameTwo, nameThree)

	nameFour := "7uckingMad!"

	fmt.Println(name, nameTwo, nameThree, nameFour)

	//Integers

	var ageOne int = 20
	var ageTwo = 30
	var ageThree int
	ageFour := 40

	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	//bits and memory

	var bitOne int8 = -124
	var bitTwo uint8 = 123

	fmt.Println(bitOne, bitTwo)

	//Float

	var floatOne float32 = 342.34
	var floatTwo float64 = 238428348.234234

	fmt.Println(floatOne, floatTwo)

}