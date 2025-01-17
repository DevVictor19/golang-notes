package main

import "fmt"

// defer -> faz com que o trecho seja executado antes da função retornar

func main() {
	defer fmt.Println("Executes before fn return")

	fmt.Println("Test!")
}
