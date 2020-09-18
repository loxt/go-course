package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println(f)

	n2 := loopFac(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func loopFac(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
