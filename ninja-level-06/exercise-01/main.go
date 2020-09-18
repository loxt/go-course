package main

import "fmt"

func main() {
	f := foo()
	b, b2 := bar()
	fmt.Println(f, b, b2)
}

func foo() int {
	return 45
}

func bar() (int, string) {
	return 2003, "Hello Loxt!"
}
