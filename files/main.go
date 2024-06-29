package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//! also done the func
	// fulfile("hello how are youuuuuuuu", "./myF.txt")
	fmt.Println("Welcome to files in golang")
	content := "whaty are happening now"
	file, err := os.Create("./myFile.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is", length)
	defer file.Close()
	readfile("./myFile.txt")

}
func readfile(filename string) {

	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("content is", string(databyte))

}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func fulfile(content string, filename string) {
	file, err := os.Create(filename)
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is", length)
	defer file.Close()
	readfile(filename)
}
