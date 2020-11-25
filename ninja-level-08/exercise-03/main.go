package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   uint8
}

func main() {
	json2 := `[{"First":"Loxt","Age":17},{"First":"Amanda","Age":19}]`

	var users []user

	err := json.Unmarshal([]byte(json2), &users)

	if err != nil {
		fmt.Println("first err", err)
	}
	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("we did something wrong", err)
	}
}
