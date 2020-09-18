package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Im", p)
}

func main() {
	p1 := person{
		first: "Emannuel",
		last:  "Loxt",
		age:   16,
	}

	p1.speak()
}
