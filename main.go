package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//Packages STRINGS

	greeting := "Hello world!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "World!"))
	fmt.Println(strings.ToLower(greeting))
	fmt.Println(strings.Index(greeting, "wo"))
	fmt.Println(strings.Split(greeting, " "))

	//the original value is unchanged
	fmt.Println("original string value =", greeting)


	//Package : SORT

	ages := []int{12, 15, 34, 57, 65, 86, 33}

	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 15)
	fmt.Println(index)

	names := []string{"barry", "anil", "li", "niyoku"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "li"))

}