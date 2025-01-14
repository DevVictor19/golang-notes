package main

import "fmt"

func main() {
	nested := map[string]map[string]float64{
		"G": {
			"Guilherme Oliveira": 32432.34,
			"Guga Pereira":       3232423.32,
		},
		"J": {
			"José João": 3232432.32,
		},
		"P": {
			"Pedro Júnior": 3232.232,
		},
	}

	delete(nested, "P")

	for _, employees := range nested {
		for name, salary := range employees {
			fmt.Printf("Name %s - Salary: %f\n", name, salary)
		}
	}
}
