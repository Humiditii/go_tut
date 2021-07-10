package main

import "fmt"

// Embedding through composition is similar to inheritance in other programming languages

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // inheriting Animal struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// b := Bird{}
	// b.Name = "Opio"
	// b.Origin = "Africa"
	// b.SpeedKPH = 33.30
	// b.CanFly = true

	b := Bird{
		Animal: Animal{
			Name:   "Opor",
			Origin: "Ota",
		},
		SpeedKPH: 38.90,
		CanFly:   true,
	}

	fmt.Println(b)
	fmt.Println(b.Animal)
	fmt.Println(b.Name)
}
