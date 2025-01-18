package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3                  // buffer fica cheio
	ch <- 4                  // chamada bloqueia o Buffer de 3 posições (não tem espaço)
	fmt.Println("Executou!") // não executa, pois o buffer está cheio
	ch <- 5
	ch <- 6
}

// Channels de buffer enchem até o limite máximo, e caso você tente inserir
// mais dados no buffer cheio ele irá bloquear a execução da go routine
// até que alguma posição do buffer seja liberada

func main() {
	ch := make(chan int, 3) // Buffer with 3 positions
	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
