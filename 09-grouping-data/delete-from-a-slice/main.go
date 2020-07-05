package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55}
	fmt.Println(x[1])
	x = append(x, 66, 77, 88, 99)
	fmt.Println(x)

	y := []int{111, 222, 333, 444}
	x = append(x, y...)
	fmt.Println(x)

	// delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
