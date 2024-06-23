package main

import "fmt"

func main() {
	var mixarray [5]interface{}
	mixarray[0] = 1
	mixarray[1] = "hello"
	mixarray[2] = 3.14
	mixarray[3] = true
	mixarray[4] = "world"
	fmt.Println(mixarray)
}
