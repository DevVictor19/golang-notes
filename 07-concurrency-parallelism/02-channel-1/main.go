package main

import "fmt"

// channel is a type
// O channel é usado para trafegar dados entre go routines
// Usado para sincronizar código assíncrono
// O código só executa quando o dado esperado chegar no canal

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // rebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
