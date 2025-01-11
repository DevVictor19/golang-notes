package main

import "fmt"

func main() {
	fmt.Print("Same ")
	fmt.Print("Line")

	fmt.Println("Create new line")
	fmt.Println("Create new line")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("Value of x is " + xs)
	fmt.Println("Value of x is", x)

	fmt.Printf("Value of x is %.2f \n", x)

	a := 1
	b := 1.9999
	c := false
	d := "Ops"

	fmt.Printf("%d %f %.1f %t %s \n", a, b, b, c, d)
	fmt.Printf("%v %v %v %v %v", a, b, b, c, d)
}
