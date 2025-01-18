package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

// Uma go routine é como se fosse uma thread, porém mais leve
// Uma go routine só continua sua execução se a thread principal estiver executando ainda

func main() {
	// go speak("Julia", "Why'd you only call me when you're high?", 500)
	// go speak("Antonio", "Are you there?", 500)
	// time.Sleep(time.Second * 10)

	go speak("Julia", "Why'd you only call me when you're high?", 500)
	speak("Antonio", "Are you there?", 10)
}
