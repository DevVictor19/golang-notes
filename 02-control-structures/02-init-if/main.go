package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s).Intn(10)
}

func main() {
	if num := randomNumber(); num > 5 {
		fmt.Println("You Win, number:", num)
	} else {
		fmt.Println("You lose, number:", num)
	}
}
