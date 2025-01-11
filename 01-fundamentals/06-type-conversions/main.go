package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int32(grade)
	fmt.Println(finalGrade)

	// convert to unicode representation (rune type)
	fmt.Println("Unicode of 97 is: " + string(97))

	// int to string
	fmt.Println("Test " + strconv.Itoa(32))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println("Number:", num)

	// string to bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Is true!")
	}
}
