package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)

	sa1 := secretAgent{person{first: "loxt", last: "god", age: 3}, "something collision", true}

	fmt.Println(sa1.person.first)
	fmt.Println(sa1.first)
}
