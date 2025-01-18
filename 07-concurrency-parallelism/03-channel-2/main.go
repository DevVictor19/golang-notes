package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines (ponto de sincronismo)
// Channel é um tipo como qualquer outro tipo da linguagem

func multiply(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)

	go multiply(2, c)
	fmt.Println("A")

	// Enquanto o dado não é lido do canal ele para execução da função
	// Espera ficar pronto o dado no canal para acessar e continuar a execução
	a, b := <-c, <-c
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
	fmt.Println("C")

	// todas a goroutines foram finalizadas, esse valor nunca vai vim (deadlock)
	fmt.Println(<-c)
}
