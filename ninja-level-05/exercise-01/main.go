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
		first:      "LOXT",
		last:       "GOD",
		favFlavors: []string{"Strawberry", "vanilla"},
	}

	fmt.Println(p1, p2)

	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
