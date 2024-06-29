package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Welcome to post request in golang")

	postrequest()
}

func postrequest() {
	const myurl = "http://127.0.0.1:8000/post"
	requestbody := strings.NewReader(`
	{
		"coursename":"golang",
		"noofstudents": "354",
		"masternme" : "Govind"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestbody)
	checknilerror(err)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	checknilerror(err)

	fmt.Println(string(databyte))
}
func checknilerror(err error) {
	if err != nil {
		panic(err)
	}
}
