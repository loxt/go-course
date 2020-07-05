package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[:3])
	fmt.Println(x[3:])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
