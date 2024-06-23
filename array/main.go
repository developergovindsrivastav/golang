package main

import "fmt"

func main() {
	fmt.Println("welcome to array")
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "mango"
	fruitlist[2] = "banana"
	fruitlist[3] = "orange"
	fmt.Println("list of fruit", fruitlist)
	fmt.Println("list of fruit", fruitlist[1])
	fmt.Println("list of fruit", fruitlist[2])
	fmt.Println("list of fruit", fruitlist[1:3])
	var activities = [4]string{"eat", "sleep", "code", "repeat"}
	fmt.Println("list of activities", activities)
}
