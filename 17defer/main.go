package main

import "fmt"

func main() {
	fmt.Println("Defer at last")

	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("Go")
	myDefer()
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
