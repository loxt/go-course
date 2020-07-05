package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[:4])
	fmt.Println(x[4:8])
	fmt.Println(x[8:12])
	fmt.Println(x[12:16])
}
