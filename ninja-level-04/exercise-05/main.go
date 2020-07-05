package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Println(x)

	x = append(x, x...)
	fmt.Println(x)

	x = append(x[:9], x[13:]...)
	fmt.Println(x)
}
