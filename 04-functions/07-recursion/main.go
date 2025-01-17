package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Invalid number %d", n)
	case n == 0:
		return 1, nil
	default:
		prev, _ := factorial(n - 1)
		return n * prev, nil
	}
}

func main() {
	fmt.Println(factorial(4))
}
