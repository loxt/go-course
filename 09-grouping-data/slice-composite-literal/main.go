package main

import "fmt"

func main() {
	// x:= type{values} // composite literal
	x := [5]int{1, 2, 3, 4, 5}

	fmt.Println(x)
}

// a SLICE Allows you to group together VALUES of the same TYPE
