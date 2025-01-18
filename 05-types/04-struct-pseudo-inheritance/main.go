package main

import "fmt"

type Car struct {
	name     string
	velocity int
}

type Ferrari struct {
	Car
	turboMode bool
}

func main() {
	f := Ferrari{
		Car:       Car{name: "Ferrari", velocity: 220},
		turboMode: false,
	}

	fmt.Println(f)
	fmt.Println(f.name, f.turboMode)
}
