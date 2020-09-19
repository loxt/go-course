package main

import "fmt"

func main() {
	a := foo()
	a()
	fmt.Println(a())
	a()
	fmt.Println(a())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
