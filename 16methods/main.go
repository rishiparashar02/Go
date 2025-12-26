package main

import "fmt"

func main() {
	fmt.Println("Structs")

	Rishi := User{"Rishi", 16, "India"}
	fmt.Println(Rishi)
	fmt.Printf("More specific details %+v \n", Rishi)

	Rishi.GetAge()
	Rishi.GetName()

}

type User struct {
	Name   string
	Age    int
	Addr   string
}

func (u User)GetAge() {
	fmt.Println("Age is:",u.Age)
}

func (u User)GetName() {
	u.Name = "Raj"
	fmt.Println("New name is:",u.Name)
}