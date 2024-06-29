package main

import "fmt"

func main() {
	fmt.Println("welcome to slice")

	var fruitlist = []string{"apple", "mango", "banana", "orange", "grapes"}
	var index int = 2
	var newfruitlist = append(fruitlist[:index], fruitlist[index+1:]...)
	fmt.Println(newfruitlist)

}
