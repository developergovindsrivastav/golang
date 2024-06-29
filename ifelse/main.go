package main

import "fmt"

func main() {
	number := 10
	if number == 10 {
		fmt.Println("number is 10")
	} else if number == 20 {
		fmt.Println("number is 20")
	} else {
		fmt.Println("number is not 10 or 20")
	}
	if num := 3; num == 3 {
		fmt.Println("number is 3")
	} else {
		fmt.Println("number is not 3")
	}

}
