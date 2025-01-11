package main

import (
	"fmt"
	m "math" // named import
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2 // type float64 inferred by go compiler

	// create new variable
	// most used way
	area := PI * m.Pow(radius, 2)
	fmt.Println("Circumference area:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var name1, name2 = "Antonio", "Victor"
	fmt.Println(name1, name2)

	color1, number := "Orange", 100
	fmt.Println(color1, number)
}
