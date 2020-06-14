package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("%T ", a)
	fmt.Println(a)
	fmt.Printf("%T ", b)
	fmt.Println(b)

	// Converted
	a = int(b)
	fmt.Printf("%T ", a)
	fmt.Println(a)
}
