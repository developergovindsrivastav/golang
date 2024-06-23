package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Current time is", time.Now())
	currenttime := time.Now()
	formatted := currenttime.Format("02-01-2006")
	fmt.Println(formatted)
}
