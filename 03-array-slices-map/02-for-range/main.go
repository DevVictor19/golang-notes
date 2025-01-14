package main

import "fmt"

func main() {
	// [...] -> faz o compilador contar pra você o tamanho fixo do array
	numbers := [...]int{1, 2, 3, 4, 5, 6}

	for index, value := range numbers {
		fmt.Printf("index: %d => value: %d\n", index, value)
	}

	for _, value := range numbers {
		fmt.Printf("value: %d\n", value)
	}
}
