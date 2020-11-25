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
	u1 := user{
		First: "M",
		Age:   54,
	}

	u2 := user{
		First: "M",
		Age:   54,
	}

	res, err := json.Marshal([]user{u1, u2})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}
