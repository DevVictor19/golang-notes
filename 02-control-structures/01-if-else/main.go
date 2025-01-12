package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Aprovado")
	} else if grade >= 5 {
		fmt.Println("Prova final")
	} else {
		fmt.Println("Rprovado")
	}
}

func main() {
	printResult(7)
	printResult(6.9)
	printResult(5)
}
