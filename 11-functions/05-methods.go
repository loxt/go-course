package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person person
	ltk    bool
}

// func (r receiver) identifier(parameters) (return(s)) { ...code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.person.first, s.person.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Emannuel",
			last:  "Loxt",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}
