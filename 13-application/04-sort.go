package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 4, 1, 3, 2, 4, 6, 7, 5, 8, 9, 10, 1, 2, 100, 25}
	sort.Ints(xi)

	fmt.Println(xi)

	sort.Sort(sort.Reverse(sort.IntSlice(xi)))

	fmt.Println(xi)
}
