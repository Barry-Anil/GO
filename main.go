package main

import (
	"fmt"
	"strings"
)

func getInitials(s string) (string, string) {
	n := strings.ToUpper(s)
	names := strings.Split(n, " ")

	var initials []string
	for _,v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	p, i := getInitials("Barry Anil")
	fmt.Printf("first name %v and Last name %v", p, i)
}
