package main

import "fmt"

func calcAverage(arr *[3]float64) float64 {
	sum := 0.0
	for i := 0; i < len(*arr); i++ {
		sum += (*arr)[i]
	}
	return sum / float64(len(*arr))
}

func printGradeAverage(average float64) {
	fmt.Printf("Grades average: %.2f\n", average)
}

func main() {
	// possuí tamanho estático
	// contém items do mesmo tipo

	var grades [3]float64

	// inicializado com 3 posições
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1

	grades2 := [3]float64{7.8, 4.3, 9.1}

	printGradeAverage(calcAverage(&grades))
	printGradeAverage(calcAverage(&grades2))
}
