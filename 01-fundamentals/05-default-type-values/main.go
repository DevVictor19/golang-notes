package main

import "fmt"

func main() {
	var a int     // 0
	var b float64 // 0
	var c bool    // false
	var d string  // empty string ""
	var e *int    // nil

	fmt.Printf("%v %v %v %d %v", a, b, c, d, e)
}
