package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a, &a)

	var b *int = &a
	fmt.Println(b, *b)

	c := &b
	fmt.Println(c, *&c, **c)

	*b = 2
	fmt.Println(a)
}
