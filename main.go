package main

import (
	"fmt"
	"math"
)

func  greeting(n string){
	fmt.Println("Ni hao",n)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func CircleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	text := "Hello"
	greeting(text)

	names := []string{"barry", "anil", "lee", "hayashi"}

	cycleNames(names, greeting)

	a1 := CircleArea(3)
	fmt.Printf("Area of the circle %0.3f",a1)

}
