package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5)
	fmt.Println("the total is", x)
}

// func (r Receiver) identifier(parameter(s)) (return(s)) { ... }

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total which is now", sum)
	}
	return sum
}
