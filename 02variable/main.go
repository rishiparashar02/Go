package main

import "fmt"

const PI string = "threepointonefour" //public

func main()  {
	fmt.Println("oii")

	var x int = 52
	fmt.Println(x)
	fmt.Printf("Variabel: %T \n", x)

	var name string = "Rishi"
	fmt.Println(name)
	fmt.Printf("Variabel: %T \n", name)

	var married bool = false
	fmt.Println(married)
	fmt.Printf("Variabel: %T \n", married)

	fmt.Println(PI)
	fmt.Printf("Constant: %T \n", PI)
}


