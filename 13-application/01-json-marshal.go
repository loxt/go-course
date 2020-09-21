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
	p1 := person{
		First: "Emannuel",
		Last:  "Loxt",
		Age:   16,
	}

	p2 := person{
		First: "Sara",
		Last:  "Tancredi",
		Age:   18,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(&people)

	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(string(bs))
}
