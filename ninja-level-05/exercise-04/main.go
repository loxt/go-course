package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:   "James",
		friends: map[string]int{"Moneypenny": 555, "Loxt": 16, "test": 3},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
}
