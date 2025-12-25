package main

import "fmt"

func main() {
	fmt.Println("Structs")

	Rishi := User{"Rishi", 16, "India"}
	fmt.Println(Rishi)
	fmt.Printf("More specific details %+v", Rishi)

}

type User struct {
	Name   string
	Age    int
	Addr   string
}