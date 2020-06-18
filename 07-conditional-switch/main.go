package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (30 == 30):
		fmt.Println("prints")
		fallthrough
	case (30 == 31):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true2")
		fallthrough
	case ("teste" == "teste"):
		fmt.Println("also true, does it print?")
		fallthrough
	default:
		fmt.Println("this is default")
	}
	n := "Bond"

	//switch "Bond" {
	switch n {
	case "Moneypenny", "Bond", "dr no":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("this is m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")

	}
}
