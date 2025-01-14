package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	s1 := []int{}

	// arr1 := append(arr1, 1, 2, 3)
	s1 = append(s1, 4, 5, 6)

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(arr1)

	s2 := make([]int, 2)

	fmt.Println(copy(s2, s1))
	fmt.Println(s2)
}
