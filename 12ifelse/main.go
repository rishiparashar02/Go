package main

import "fmt"

func main() {
	fmt.Println("if else")
	age := 17

	if age < 18 {
		fmt.Println("You are a minor")
	} else if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are 18 years old")
	}

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	if num:= 5; num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is non-negative")
	}
	
}
