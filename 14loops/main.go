package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// Traditional for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//similar to for each
	for i:= range days {
		fmt.Println(days[i])
	}


	for index, day := range days {
		fmt.Printf("Index is %v and day is %v\n", index, day)
	}

	roughvalue := 1

	for roughvalue < 10 {

		if roughvalue == 5 {
			break
		}

		if roughvalue == 2{
			goto l
		}
		fmt.Println("value is" , roughvalue)
		roughvalue++
	}

l:
	fmt.Println("Jumped to label l")
}
