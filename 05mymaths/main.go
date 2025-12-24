package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome")

	// var num1 int = 1
	// var num2 float64 = 4.551

	// fmt.Println(num1 + int(num2))


	//random numbers

	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(6) + 2)


	// randomness by crypto
	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)
}
