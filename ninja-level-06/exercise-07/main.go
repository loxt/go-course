package main

import "fmt"

func main() {
	an := func() {
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
		fmt.Println("done")
	}
	an()

}
