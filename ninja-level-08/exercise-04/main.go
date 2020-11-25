package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 5, 3200, 15, 2, 1, 2, 3, 10, 230, 512, 1024}
	xs := []string{"loxt", "brabo", "monstro", "asefude"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
}
