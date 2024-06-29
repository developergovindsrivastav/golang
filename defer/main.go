package main

//! defer executes at the last of the function

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") //* last defer is always execute first..
	fmt.Println("hello!")
	print()
}
func print() {

	for i := 0; i < 5; i++ {
		defer fmt.Println("the no is", i) //? due to the defer the digits are inversing
	}
}
