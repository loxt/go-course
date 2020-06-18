package main

import (
	"fmt"
)

func main() {
	if x := 42; x == 42 {
		fmt.Println("A short function return true")
	}

	if true {
		fmt.Println("Hello, world!")
	}
	if false {
		fmt.Println("Hello, world but with false!")
	}

	if !true {
		fmt.Println("Hello, world but without true !")
	}

	if !false {
		fmt.Println("Hello, world but without false !")
	}

	if 2 == 2 {
		fmt.Println("Hello, world but 2 is equal 2!")
	}

	if 2 != 2 {
		fmt.Println("Hello, world but 2 isn't equal 2!")
	}

	if !(2 == 2) {
		fmt.Println("Hello, world but 2 is equal 2!")
	}

	if !(2 != 2) {
		fmt.Println("Hello, world but 2 isn't equal 2!")
	}
}
