package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) changeMyName(name string) {
	p.first = name
}

func main() {
	p1 := person{
		first: "Emannuel",
		last:  "Loxt",
	}
	fmt.Println(p1)
	p1.changeMyName("Loxtao")
	fmt.Println(p1)
}
