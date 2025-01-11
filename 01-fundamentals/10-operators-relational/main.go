package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">", 3 > 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("d1 == d2", d1 == d2) // compare by value not reference
	fmt.Println("d1.Equal(d2)", d1.Equal(d2))

	type Person struct {
		Nome string
	}

	p1 := Person{"Antonio"}
	p2 := Person{"Victor"}
	fmt.Println("Persons", p1 == p2) // compare by value not reference
}
