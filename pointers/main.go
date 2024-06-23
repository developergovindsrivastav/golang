package main

import "fmt"

func main() {
	fmt.Println("welcome in pointer")
	myno := 23

	var ptr = &myno

	fmt.Println("value of myno is", myno)
	fmt.Println("value of myno is", *ptr)
	fmt.Println("value of myno is", &ptr)
	*ptr = *ptr + 1

	fmt.Println("value of myno is", myno)

}
