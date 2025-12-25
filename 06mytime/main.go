package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome 6")

	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	cratedate := time.Date(2020, time.March, 25, 10, 30, 0, 0, time.UTC)
	fmt.Println(cratedate)
	fmt.Println(cratedate.Format("01-02-2006 Monday"))
}
