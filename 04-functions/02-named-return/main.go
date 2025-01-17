package main

import "fmt"

func swipe(first, second int) (s, f int) {
	f = first
	s = second
	return // clean return (s, f) are internal variables for return
}

func main() {
	s, f := swipe(10, 5)
	fmt.Println(s, f)
}
