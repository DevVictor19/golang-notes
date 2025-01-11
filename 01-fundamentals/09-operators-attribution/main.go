package main

import "fmt"

func main() {
	var b uint = 3
	fmt.Println(b)

	i := 3 // type inference
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}
