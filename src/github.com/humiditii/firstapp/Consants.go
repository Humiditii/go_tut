package main

import "fmt"

const (
	a = 1 << iota
)

func main() {
	fmt.Println(a)

	fmt.Println("Constants")
	const a int = 24
	const b int = 43
	fmt.Println(a + b)
}
