package main

import "fmt"

func main() {
	coord1 := Coord{2.0, 2.0}
	coord2 := Coord{2.0, 4.0}

	fmt.Println(legs(coord1, coord2)) // same package
	fmt.Println(Distance(coord1, coord2))
}
