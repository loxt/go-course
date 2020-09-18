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
	fmt.Println("I am", s.person.first, s.person.last, " - the secretAgent speak ")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak ")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
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

	p1 := person{
		first: "Dr.",
		last:  "Loxt",
	}

	sa1.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(p1)
}
