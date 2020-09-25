package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   uint8
}

func main() {
	s := `[{"First":"Emannuel","Last":"Loxt","Age":16},{"First":"Sara","Last":"Tancredi","Age":18}]`

	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("people number", i, ": ", v)
	}
}
