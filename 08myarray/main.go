package main

import "fmt"

func main() {
	fmt.Println("welcome to array")

	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[2] = "grapes"
	fruitlist[3] = "mango"
	fmt.Println("fruitlist is:", fruitlist)
	fmt.Println("length of fruitlist is:", len(fruitlist))

	var veglist = [3]string{"potato", "tomato", "brinjal"}
	fmt.Println("veglist is:", veglist)
	fmt.Println("len ", len(veglist))

}
