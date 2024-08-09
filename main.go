package main

import (
	"fmt"
)



func main() {

	menu := map[string]float64{
		"soup" : 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"chocolate Pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k,v := range menu {
		fmt.Println(k, "-", v)
	}

	records := map[int]string {
		1 : "barry",
		2 : "anil",
		3 : "lee",
	}

	fmt.Println(records)

	records[1] = "Barry Anil"
	fmt.Println(records)

	records[2] = "Anil Kapoor"
	fmt.Println(records)

}
