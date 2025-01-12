package main

import (
	"fmt"
	"time"
)

func getResult(grade float64) string {
	// o switch do go executa sempre um case e sai do switch
	// para continuar para o proximo case use fallthrough
	integer := int(grade)
	switch integer {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "F"
	default:
		return "Invalid grade"
	}
}

func printGreetings() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good night!")
	}
}

func main() {
	fmt.Println(getResult(10))
	fmt.Println(getResult(9.8))
	printGreetings()
}
