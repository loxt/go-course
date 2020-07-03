package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[1] = 1
	fmt.Println(x)
	x[3] = 1
	fmt.Println(x)
	fmt.Println(len(x))
}
