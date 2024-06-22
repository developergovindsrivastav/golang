package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating: ")
	input, _ := reader.ReadString('\n')
	ratingno, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("your rating is", ratingno+2)
	}

}
