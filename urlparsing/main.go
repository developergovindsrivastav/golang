package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("welcome to the url parsing in golang")

	const Url string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjk"
	result, _ := url.Parse(Url)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	queries := result.Query()
	fmt.Printf("type of queries is %T\n", queries)
	for _, val := range queries {
		fmt.Println(val)
	}
	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/courses",
		RawQuery: "coursename=golang&paymentid=56545465",
	}
	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)

}
