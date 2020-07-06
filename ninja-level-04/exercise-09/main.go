package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":         []string{"James bond", "James", "Bond"},
		"Shaken not stirred": []string{"Shaken", "Sha", "Ken"},
		"Martinis":           []string{"It's Martinis", "Its", "Mart"},
		"Women":              []string{"That's a Women", "Thats", "Wom"},
	}
	fmt.Println(m)

	m["Loxt"] = []string{"steaks", "cigars"}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)

		}
	}

}
