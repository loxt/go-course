package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	foo()
	bar("loxt")
	s1 := woo("Loxt")
	fmt.Println(s1)
	x, y := mouse("Emannuel", "Loxt")
	fmt.Println(x, y)
}

//func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("hello", s+"!")
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "Hello": `)
	//b := false
	return a, false
}
