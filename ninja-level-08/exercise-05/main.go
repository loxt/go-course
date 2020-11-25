package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   uint8
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByFirst []user

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	json2 := `[{"First":"Loxt","Age":17},{"First":"Amanda","Age":19}]`

	var users []user

	err := json.Unmarshal([]byte(json2), &users)

	if err != nil {
		fmt.Println("first err", err)
	}

	for _, v := range users {
		fmt.Println("Unsorted:", v)
	}

	sort.Sort(ByFirst(users))
	for _, v := range users {
		fmt.Println("Sorted by First", v)
	}

	sort.Sort(ByAge(users))
	for _, v := range users {
		fmt.Println("Sorted by Age", v)
	}
}
