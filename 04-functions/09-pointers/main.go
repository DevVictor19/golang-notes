package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// acessando valor armazenado no ponteiro
	*n++
}

func modifyMyMap(m map[int]string) {
	m[2] = "Modified"
}

// No golang sempre estamos trabalhando com uma cópia do valor
// A não ser que vocẽ explicitamente use uma referência
// Com exceção de slices e maps que são sempre por referência
func main() {
	n := 1

	// passando uma cópia do valor
	inc1(n)

	// paassando a referência de um valor
	inc2(&n)

	fmt.Println(n)

	// map por padrão usa referência
	m := map[int]string{
		1: "First",
		2: "Second",
	}

	modifyMyMap(m)

	fmt.Println(m)
}
