package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer fmt.Println("foo")
}

func bar() {
	defer fmt.Println("bar")
}
