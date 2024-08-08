package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("the value of x :", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("the value of i :", i)
	}

	names := []string{"barry", "anil", "lee", "mario"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)

	}
}
