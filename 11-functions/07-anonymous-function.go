package main

import "fmt"

func main() {
	rand()

	func(x int) {
		fmt.Println("Anonymous func")
		fmt.Println("Thea meaning of life", x)
	}(42)
}

func rand() {
	fmt.Println("foo ran")
}
