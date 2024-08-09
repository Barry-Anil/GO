package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	initialBill := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "choco lava": 6.99},
		tip:   0,
	}

	return initialBill
}

func (b bill) format() string {

	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ..... %v \n", k+":", v)
		total += v
	}

	//total 
	fs += fmt.Sprintf("%-25v ..... $%0.2f", "total:", total)

	return fs
}