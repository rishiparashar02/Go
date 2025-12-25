package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")
	var fruitlist = []string{"apple", "banana", "grapes"}
	fmt.Println("fruitlist is:", fruitlist)
	fmt.Printf("fruitlist is: %T", fruitlist)

	var veglist = []string{"potato", "tomato", "brinjal"}
	fmt.Println("veglist is:", veglist)

	veglist = append(veglist, "carrot", "beans")
	fmt.Println("new veglist is:", veglist)
	fmt.Println("new veglist is:", veglist[:3])

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 6556
	highscores[2] = 178
	highscores[3] = 890

	highscores = append(highscores, 4567, 9876) // this automatically changes the size of slice
	fmt.Println("highscores:", highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println("sorted highscores:", highscores)

	var courses = []string{"reactjs", "javascript", "html", "css", "vuejs", "docker"}
	fmt.Println("courses:", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses after deletion:", courses)

}
