package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	var ptr *int
	fmt.Println("value of ptr is", ptr)

	var num = 21
	var ptrr = &num
	fmt.Println("address of num is", ptrr)
	fmt.Println("value at address ptrr is", *ptrr)

	*ptrr = *ptrr + 10
	fmt.Println("new value:" , num)

}
