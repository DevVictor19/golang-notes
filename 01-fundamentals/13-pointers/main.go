package main

import "fmt"

func main() {
	num := 100

	// Go não tem aritmética de ponteiros (operações com pointeiros)
	// Aloco um espaço para armazenar um ponteiro
	// Nesse espaço de memória eu armazeno o endereço de memória de outro dado
	var pointer *int = &num
	fmt.Println("Acessar endereço de memória da variável:", &num)

	fmt.Println("Acessar endereço de memória armazenado no ponteiro:", pointer)

	*pointer++
	fmt.Println("Acessar endereço de memória do ponteiro:", &pointer)

	fmt.Println("Acessar valor armazenado no endereço de memória pelo ponteiro:", *pointer)
	fmt.Println("Acessar valor da variável original:", num)
}
