package main

import "fmt"

func mofifyMap(m map[string]int) {
	m["newKey"] = 3
	m["key1"] = 100
}

func main() {
	originalMap := map[string]int{
		"key1": 1,
		"Key2": 2,
	}

	fmt.Println("Before modification", originalMap)

	mofifyMap(originalMap)

	fmt.Println("After mofification", originalMap)
}
