package main

import (
	"fmt"
)

func main() {
	agesMap := map[string]int{
		"Hameed":   19,
		"Hameedah": 11,
		"Azeezah":  33,
		"Halimah":  11,
	}

	fmt.Println(agesMap)

	agesMap["hafiz"] = 30
	fmt.Println(agesMap)

	school := make(map[string]int)
	fmt.Println(school)
	school = map[string]int{
		"FUTA": 200,
		"OAU":  20,
	}
	fmt.Println(school)

	// deleting content
	delete(agesMap, "Hameed")

	fmt.Println(agesMap)
}
