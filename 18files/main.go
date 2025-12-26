package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files ")
	content := "This is a sample content of the file."

	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}