package main

import "fmt"

type course struct {
	name string
}

func main() {
	// interface atua como um tipo genérico
	var thing interface{} = course{}

	fmt.Println(thing)

	type any interface{}

	var thing2 any = "Opa"

	fmt.Println(thing2)

	thing2 = 123

	fmt.Println(thing2)
}
