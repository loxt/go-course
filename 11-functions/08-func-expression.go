package main

import "fmt"

func main() {
	rand()

	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("My second func expression")
	}
	f2(42)
}
