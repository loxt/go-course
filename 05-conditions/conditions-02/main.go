package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("Dammit, x isn't equals 42, our value was 40!")
	} else if x == 41 {
		fmt.Println("Dammit, x isn't equals 42, our value was 41!")
	} else {
		fmt.Println("Success, x is equals 42")
	}
}
