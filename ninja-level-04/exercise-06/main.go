package main

import "fmt"

func main() {
	x := make([]string, 20, 50)
	fmt.Println(x)
	x = []string{"loxt", "GOD"}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
