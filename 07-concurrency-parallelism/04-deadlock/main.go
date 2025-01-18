package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante (só desbloqueia quando o dado for lido)
	fmt.Println("Só depois que for lido")
}

// Quando usamos um Buffer, ocorre o bloqueio apenas depois do Buffer ficar cheio
// Assim o desbloqueio só ocorre após alguém consumir o Buffer e liberar espaço

func main() {
	c := make(chan int) // canal sem buffer

	go routine(c)

	fmt.Println(<-c) // operação bloqueante (só desbloqueia após o dado chegar e for lido)
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock (dado nunca vai chegar e vai bloquear a thread principal)
}
