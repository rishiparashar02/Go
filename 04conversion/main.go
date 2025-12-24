package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to shop") //print statement
	fmt.Println("Rate the items")

	reader := bufio.NewReader(os.Stdin) // Creates a reader to read input from keyboard

	input, _ := reader.ReadString('\n') // Reads user input until Enter key is pressed (ignores error)

	fmt.Println("THanks for rating:", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)  // Converts input string to float after removing spaces/newline

	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("Added 1 to rating", numrating + 1)
	}
}
