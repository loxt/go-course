package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)

	p2 := person{
		first: "loxt",
		last:  "god",
		age:   27,
	}

	fmt.Println(p2)
}
