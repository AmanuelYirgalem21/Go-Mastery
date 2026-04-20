package main

import "fmt"

func main() {
	marks := map[string]int{
		"Amanuel": 90,
		"John":    75,
		"Mary":    85,
	}

	name := "John"

	// check if key exists
	if score, ok := marks[name]; ok {
		fmt.Println(name, "scored", score)
	} else {
		fmt.Println(name, "not found")
	}
}