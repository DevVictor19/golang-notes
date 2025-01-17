package main

import "fmt"

func average(grades ...float64) float64 {
	sum := 0.0

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))
}

func printNames(names ...string) {
	fmt.Println(names)
}

func main() {
	fmt.Println(average(3, 4, 3.3, 4.3, 10))

	// only works with slices!
	names := []string{"Maria", "Rafael", "João", "Júlia"}
	printNames(names...)
}
