package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello slice \n")

	// creating a slice
	a := []int{2, 4, 5}

	fmt.Printf("Slice a is :-%v, type: %T \n", a, a)

	// len function is common to both array and slice
	// cap function for the capacity
	// assigning a new variable to an already assigned slice refers to the same elements
	b := a
	b[1] = 9

	fmt.Println(b)
	fmt.Println(a)

	// slicing a slice is the same to slicing in python i.e [n:m], [n:], [:m]

	// Another method is to use to use the make function: 1st: type of object, 2nd:length 3rd: capacity

	c := make([]int, 4, 20)

	fmt.Println(c)
	fmt.Printf("Length : %v \n", len(c))
	fmt.Printf("Capacity : %v \n", cap(c))

	// adding to the slice
	c = append(c, 40)

	fmt.Printf("Length : %v \n", len(c))
	fmt.Printf("Capacity : %v \n", cap(c))

	c = append(c, 90, 58, 6, 8)
	fmt.Printf("Length : %v \n", len(c))
	fmt.Printf("Capacity : %v \n", cap(c))

	// concatenating slices
	// slice.append(slice, []int{a,b,c}...)  // just like JS spread operator
	// poping
	a = a[1:]
	fmt.Println(a)
}
