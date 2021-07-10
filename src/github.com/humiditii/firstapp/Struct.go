package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Azeezah",
		companions: []string{
			"Azeezah",
			"KUnle",
			"Uthman",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	// Annonymous struct
	myProfile := struct {
		name   string
		status string
	}{name: "Hameed", status: "True"}

	// NOte if we want to point to same data we use & operator

	fmt.Println(myProfile)

}
