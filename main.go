package main

import "fmt"

func main() {
	age := 29
	name := "barry"
	//Print 
	fmt.Print("hello")
	fmt.Print("hello \n")
	fmt.Print("hello \n")

	//Print line
	fmt.Println("Hello, Ninjas!!")

	fmt.Println("My name is =", name, "age is =", age)

	//priintf (formatted Strings) %_ = format specifier
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("my name is %q and my age is %q \n", name, age)
	fmt.Printf(" age is of Type %T \n", age)
	fmt.Printf("your score is %f \n", 98.99)
	fmt.Printf("your score is %0.1f \n", 98.9924)


	// SprintF (save formatted strings)
	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	fmt.Println(str)
}