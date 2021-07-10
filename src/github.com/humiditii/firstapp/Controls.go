package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The statement is true")
	}

	testMap := map[string]int{
		"ota":   20,
		"sango": 15,
		"iyana": 40,
	}

	// initializer
	if present, ok := testMap["ota"]; ok {
		fmt.Println(present)
	}

	if testMap["ota"] > testMap["sango"] {
		fmt.Println(" Ota is more than Sango")
	}

	if testMap["sango"] != testMap["iyana"] {
		fmt.Println("sango not equal to iyana")
	}
}
