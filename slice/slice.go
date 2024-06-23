package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice")
	var fruitlist = []string{"apple", "mango", "banana", "orange"}
	fmt.Println("list of fruit", fruitlist)
	fruitlist = append(fruitlist, "kiwi")
	fmt.Println("list of fruit", fruitlist)
	fmt.Println("slice", fruitlist[1:3]) //! 3 is not includes
	fmt.Println("slice", fruitlist[:3])  //! Starts from 0
	fmt.Println("slice", fruitlist[2:])  //! starts from 2 and ends on the last
	//? best facts-----
	//! we can also s=do than we can add the more elements without seeing the list size
	var prabhuramkenaam = []string{"dashrathnandan", "kausalyanandan", "ruaghukulnayak", "aaryasarvoach", "raghukulsiromani"}
	prabhuramkenaam = append(prabhuramkenaam, "prabu")
	fmt.Println(prabhuramkenaam)

	//! Sorting the array
	sort.Strings(prabhuramkenaam)
	fmt.Println(prabhuramkenaam)
}
