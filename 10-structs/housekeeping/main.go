package main

import "fmt"

// we create VALUES of a certain TYPE that are stored in VARIABLES.
// and those VARIABLES have identifiers.

var x int

type person struct {
	first string
	last  string
}

type foo int

var y foo

//const bar int = 42
const bar = 42

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1)

	y = 42
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)

	y = bar
	fmt.Println(y)
}
