package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   uint8
}

func main() {
	json2 := `[{"First":"Loxt","Age":17},{"First":"Amanda","Age":19}]`

	var people []user

	err := json.Unmarshal([]byte(json2), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Print("User number #", i, " have the current name: ", v.First+"\n")
	}
}
