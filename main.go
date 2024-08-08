package main

import "fmt"

func main() {
	
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"barry", "anil", "lee", "nihayo"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at Pos: ", index)
			continue
		}
		if index > 2 {
			fmt.Println("break at pos", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
