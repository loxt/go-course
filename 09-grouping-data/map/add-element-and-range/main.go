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
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["Loxt"] = 16
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
