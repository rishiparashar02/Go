package main

import "fmt"

func main() {
	fmt.Println("maps aagaya")

	languages := make(map[string] string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"
	fmt.Println("languages:", languages)

	fmt.Println("js shorts for", languages["js"])
	delete(languages, "rb")
	fmt.Println("languages:", languages)
	fmt.Println("length is:", len(languages))

	// loops in go

	for key, value := range languages { //comms, ok
		fmt.Printf("for key %s value is %s\n", key, value)
    }
}

