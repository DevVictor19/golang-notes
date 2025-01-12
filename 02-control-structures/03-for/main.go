package main

import "fmt"

func main() {
	// única estrutura de repetição da linguagem
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Using external variable -------")

	count := 0

	for count < 10 {
		fmt.Println(count)
		count++
	}

	condition := true

	// loop infinito
	for {

		if condition {
			fmt.Println("Breaking...")
			break
		}

	}
}
