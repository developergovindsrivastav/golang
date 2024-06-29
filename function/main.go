package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the number: ")
	no := bufio.NewReader(os.Stdin)

	input, _ := no.ReadString('\n')
	intno, _ := strconv.Atoi(strings.TrimSpace(input))

	fmt.Println(add(intno, 5))
	fmt.Println(manyadd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func add(x, y int) int {
	return x + y
}
func manyadd(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "hello"
}
