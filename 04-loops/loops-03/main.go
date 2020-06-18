package main

import "fmt"

func main() {
	i := 1
	//for i < 2 {
	//	i++
	//	fmt.Println("Hello world!")
	//}

	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("Done.")
}
