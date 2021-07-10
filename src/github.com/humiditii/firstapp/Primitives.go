package main

import (
	"fmt"
)

func main() {
	// boolean
	var n bool = true
	// float
	var a float32 = 3.14
	a = 33e12
	a = 2.1e20
	fmt.Printf("%v,%T\n", n, n)
	fmt.Println(a)
	// NUmeric types

	// types: int, int8, int16, int32, int64, uint8, uint16, uint32

	//  complex numbers
	var current complex128 = 2 + 4i
	var voltage = complex(2, 5)

	fmt.Printf("%v, %T", imag(current), imag(current))
	fmt.Printf("%v, %T", real(current), real(current))
	fmt.Printf("%v, %T\n", real(voltage), real(voltage))

	s1 := "My name is "
	s2 := " Hameed Babatunde"
	s3 := "Harleemah"

	fmt.Println([]byte(s3))

	fmt.Println(s1 + s2)
}
