package main

import "fmt"

func main() {
	defer foo2()
	bar2()
}

func foo2() {
	defer fmt.Println("foo")
}

func bar2() {
	defer fmt.Println("bar")
}
