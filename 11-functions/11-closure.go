package main

import "fmt"

func main() {
	a := incrementor()
	a()
	a()
	a()
	a()
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		fmt.Println(x)
		return x
	}
}
