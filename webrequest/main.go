package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {

	fmt.Println("Hello, WebRequest")
	response, err := http.Get(url)
	checknilError(err)

	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	checknilError(err)
	conent := string(databyte)

	fmt.Println(conent)

}
func checknilError(err error) {
	if err != nil {
		panic(err)
	}
}
