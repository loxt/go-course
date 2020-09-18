package main

import "fmt"

func main() {
	s1 := returnF()
	fmt.Println(s1)

	s2 := bar()

	fmt.Printf("%T\n", s2)
}

func returnF() string {
	return "Hello World"
}

func bar() func() int {
	return func() int {
		return 2003
	}
}
