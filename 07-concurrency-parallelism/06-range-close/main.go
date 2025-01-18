package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int, c chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				c <- prime
				start = prime + 1
				time.Sleep(time.Second)
				break
			}
		}
	}
	close(c) // closing channel
}

func main() {
	c := make(chan int, 30)
	go primes(cap(c), c)

	// sempre que chegar um valor o laço vai iterar...
	for prime := range c {
		fmt.Printf("%d ", prime)
	}
}
