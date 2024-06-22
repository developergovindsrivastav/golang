package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var a int = 50
	fmt.Printf("Value of a is %T,\n", a)
	var b float64 = 50.23333333333
	fmt.Printf("Value of b is %T,\n", b)
	var c bool = true
	fmt.Printf("Value of c is %T,\n", c)
	var d string = "Hello"
	fmt.Printf("Value of d is %T,\n", d)

	fmt.Println(a, b, c, d)
}
