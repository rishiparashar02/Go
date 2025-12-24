package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wel := "welcome"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	//comma ok || err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks 4 rating" , input)
	fmt.Printf("Type of input is %T \n", input)
}