package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func exec(cb func(a, b int) int, p1, p2 int) int {
	return cb(p1, p2)
}

func main() {
	fmt.Println(exec(multiply, 1, 2))
}
