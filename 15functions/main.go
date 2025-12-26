package main

import "fmt"

func main() {
	fmt.Println("functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result from adder function:", result)

	proresult, _ := proadder(2, 4, 6, 8, 10)
	fmt.Println("Result from proadder function:", proresult)	
}
 
func greeter()  {
	fmt.Println("Hello from greeter function")
}

func adder(num1 int, num2 int) int { // int outside the return type of result
	return num1 + num2
}

func proadder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}	
	return total, "hello"
}