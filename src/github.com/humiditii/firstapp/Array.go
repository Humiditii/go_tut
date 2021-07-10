package main

import (
	"fmt"
)

func main() {
	// inatializing array with a specific number of elements it can hold
	grades := [3]int{30, 60, 70}
	fmt.Printf(" Grades: %v \n", grades)

	// array with any size
	courses := [...]string{"MME201", "MME401"}
	fmt.Printf("Courses: %v \n", courses)

	// initializing an empty array
	var students [3]string
	fmt.Println(students)
	students[0] = "Hameed"
	students[1] = "Kunle"
	students[2] = "Lola"

	fmt.Printf("Students are: %v", students)

	fmt.Println(len(students))

	var identityMatrix [3][3]int //= [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// OR
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)

	a := []int{2, 3, 4, 5, 6}
	b := a[:2]
	c := append(a, 30)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
