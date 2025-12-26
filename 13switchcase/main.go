package main

import (
	"math/rand"
	"fmt"
	"time" 
)

func main() {
	fmt.Println("switchcase")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1, you can go out")
	case 2:
		fmt.Println("You got 2, you can go out")
	case 3:
		fmt.Println("You got 3, you can go out")
		fallthrough
	case 4:
		fmt.Println("You got 4, you can go out")
		fallthrough
	case 5:
		fmt.Println("You got 5, you can go out")
	}
}
