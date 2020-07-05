package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55}
	fmt.Println(x[1])
	x = append(x, 66, 77, 88, 99)
	fmt.Println(x)

	y := []int{111, 222, 333, 444}
	fmt.Println(append(x, y...))
}
