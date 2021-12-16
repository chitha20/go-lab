package main

import (
	"fmt"
)

func main() {
	// int initialize ( not recommended)
	var i int
	i = 10
	fmt.Println(i)

	// String initialize ( Not recommended)
	var c string
	c = "Chirag"
	fmt.Println(c)

	// Float datatype
	f := 2.14
	fmt.Println(f)

	// String assignment
	firstName := "Thaker"
	fmt.Println(firstName)

	// Boolean assignment
	b := true
	fmt.Println(b)

	// variable initialize and assignment
	var j int = 13
	fmt.Println(j)

	// pointers
	var lastName *string = new(string)
	*lastName = "Nirali"
	fmt.Println(*lastName)

	// Pointers addressing and referencing
	middleName := "Gautam"
	fmt.Println(middleName)

	ptr := &middleName
	fmt.Println(ptr, *ptr)

	// Constants and converting datatypes from int to float
	const t int = 3
	fmt.Println(float32(t) + 3.2)
}
