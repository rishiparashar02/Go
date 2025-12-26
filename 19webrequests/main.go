package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
 
const URL = "https://rishiparasharportfolio.vercel.app/"

func main() {
	fmt.Println("web requests")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Responce is of type %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
