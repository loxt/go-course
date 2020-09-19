package main

import "fmt"

func main() {
	myFunc := returnAFunction()
	myFunc()
}

func returnAFunction() func() {
	return func() {
		fmt.Println("My func returns another func")
	}
}
