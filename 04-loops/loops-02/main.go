package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i <= 15; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("\t\t The outer loop: %d\t The inner loop: %d\n", i, j)
		}
	}
}
