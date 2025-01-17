package main

import "fmt"

func closure() func() {
	x := 10

	return func() {
		fmt.Println(x)
	}
}

func main() {
	fn := closure()
	fn()
}
