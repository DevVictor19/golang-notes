package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array (fixed size)
	a1 := [3]int{1, 2, 3}

	// slice (dynamic size)
	s1 := []int{1, 2, 3} // creating a new array and storing in a slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	fmt.Println("--------------")

	// slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // reference to an array, does not create a new array

	fmt.Println("a2:", a2)
	fmt.Println("s2:", s2)

	s3 := a2[:2] // from start of array to index 2 - 1
	s4 := a2[1:] // from index 1 to end of array

	fmt.Println("s3:", s3)
	fmt.Println("s4", s4)

	s5 := s4[1:3]

	fmt.Println("s5:", s5)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	s5[0] = 99

	fmt.Println("s5:", s5)
	fmt.Println("a2:", a2)

	// slice is a box of contiguous pointers
}
