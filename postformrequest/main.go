package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	performpostformrequest()
}

func performpostformrequest() {
	const myurl = "http://127.0.0.1:8000/postform"
	data := url.Values{}
	data.Add("firstname", "Govind")
	data.Add("lastname", "Srivastav")
	data.Add("email", "govindsrivastav800@gmail.com")

	response, err := http.PostForm(myurl, data)
	checknilerror(err)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	checknilerror(err)

	content := string(databyte)
	fmt.Println(content)

}
func checknilerror(err error) {
	if err != nil {
		panic(err)
	}
}
