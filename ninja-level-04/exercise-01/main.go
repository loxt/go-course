package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	x[1] = 8
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
}
