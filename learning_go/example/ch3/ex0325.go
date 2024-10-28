package main

import "fmt"

func stringex() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b)
}

func main() {
	// var data []int
	data := []int{2, 4, 6, 8}
	fmt.Println(data)

	stringex()
}
