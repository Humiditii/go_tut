package main

import (
	"fmt"
	"strconv"
)

var name string = "Hameed"

var (
	age    int    = 22
	school string = "FUTA"
)

func main() {
	// var keyword then the identifier followed by the type and assignment or assign it in another line
	var i float32 = 22
	var j int
	j = 21

	k := 211

	var who string

	who = strconv.Itoa(j)

	// fmt.Println(k)
	// fmt.Println(j)
	// fmt.Println(i)
	fmt.Printf(" %v, %T\n ", i, i)
	fmt.Printf(" %v, %T\n", j, j)
	fmt.Printf(" %v, %T\n ", k, k)
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", school, school)
	fmt.Printf("%v,%T\n", who, who)

	// note: PrintLn prints the normal value passed in, while PrintF prints in a format %v=> value %T =>type

	// converting variables
	j = int(i)
	fmt.Println(j)
}
