package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	m["d"] = 4
	m["e"] = 5
	fmt.Println(m)

	delete(m, "a")

	fmt.Println(m)
	// for k, v := range m { //! regular syntax
	// 	fmt.Printf("key is %v and value is %v\n", k, v)
	// }
	//? comma ok syntax
	for _, value := range m {
		fmt.Printf("key is v and value is %v\n", value)
	}

}
