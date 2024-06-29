package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to get request in golang")
	getrequest()

}
func getrequest() {
	const myurl = "http://127.0.0.1:8000/get"
	response, err := http.Get(myurl)
	checknilerror(err)
	defer response.Body.Close()
	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length is", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func checknilerror(err error) {
	if err != nil {
		panic(err)
	}
}
