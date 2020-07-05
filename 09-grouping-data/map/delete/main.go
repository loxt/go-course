package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	delete(m, "No Errors")
	fmt.Println(m)
	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["No Errors"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}
