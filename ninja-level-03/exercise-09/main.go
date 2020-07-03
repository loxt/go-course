package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("should not print")
	case "surfing":
		fmt.Println("should print")
	}
}
