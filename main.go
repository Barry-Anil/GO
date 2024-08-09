package main

import (
	"fmt"
)

var score = 98.5

func main() {

	//you cant access var score = 98.5 because its outside the scope and inside main function. 

	sayHello("barry")

	for _, v := range points {
		fmt.Println(v)
	}

	yourScore()
}
