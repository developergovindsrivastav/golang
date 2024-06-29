package main

import "fmt"

func main() {
	type Person struct { //! NO inheritance , NO class parent system
		Name   string
		Age    int
		Status bool
		Email  string
	}
	govind := Person{"govind", 17, true, "govindsrivastav800@gmail.com"}
	// fmt.Printf("the keys are %+v", govind)//! the one way to print
	fmt.Printf("the name are %v and the age is %v", govind.Name, govind.Age)
}
