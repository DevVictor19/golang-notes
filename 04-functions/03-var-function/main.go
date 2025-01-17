package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

var test = func() func() {
	return func() {
		fmt.Println("Calling nested fn")
	}
}

func main() {
	fmt.Println(sum(3, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(3, 4))

	test()()
}
