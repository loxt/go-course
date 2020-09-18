package main

import "fmt"

func main() {
	foo()
	defer foo()
	fmt.Println("1")
}

func foo(variadic ...int) {
	fmt.Println("foo bar")
}
