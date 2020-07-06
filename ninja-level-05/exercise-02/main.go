package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"Chocolate", "Martini", "Rum and coke"},
	}

	p2 := person{
		first:      "loxt",
		last:       "god",
		favFlavors: []string{"Chocolate", "Martini", "Rum and coke"},
	}

	fmt.Println(p1, p2)

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
