package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("1st project")
	rand.Seed(time.Now().UnixNano())
	randomint := rand.Intn(5) + 1
	no := 0

	switch randomint {
	case 1:
		fmt.Println("your no. is 1")
		no++
		fmt.Println(no)
	case 2:
		fmt.Println("your no. is 2")
		no = +2
		fmt.Println(no)

	case 3:
		fmt.Println("your no. is 3")
		no = +3
		fmt.Println(no)

	case 4:
		fmt.Println("your no. is 4")
		no = +4
		fmt.Println(no)

	case 5:
		fmt.Println("your no. is 5")
		no = +5
		fmt.Println(no)
	case 6:
		fmt.Println("your no. is 6")
		no = +6
		fmt.Println(no)
	default:
		fmt.Println("invalid no.")

	}
	if no > 25 {
		fmt.Println("you win")

	}

}
