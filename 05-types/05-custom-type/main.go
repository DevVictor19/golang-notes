package main

import "fmt"

type Grade float64

func (g Grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func main() {
	var grade Grade = 8.0

	fmt.Println(grade.between(8, 10))
}
