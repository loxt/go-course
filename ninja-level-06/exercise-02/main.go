package main

import "fmt"

func main() {
	f := foo(1, 2, 3, 4, 5)
	cl := []int{1, 2, 3, 4, 5}
	f2 := foo(cl...)
	fmt.Println(f, f2)
}

func foo(variadic ...int) int {
	sum := 0
	for _, v := range variadic {
		sum += v
	}
	return sum
}
