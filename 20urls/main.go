package main

import "fmt"
import "net/url"

const myurl string = "https://rishiparasharportfolio.vercel.app/"

func main() {
	fmt.Println("url")
	_ = myurl

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _ ,val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "rishiparasharportfolio.vercel.app",
		Path: "/hsddfbvujrihsuofb",
		RawPath:  "user=Rishi",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("The constructed url is: ", anotherUrl)
}
