package main

import "fmt"

func main() {
	players := map[int]string{}

	players[1] = "John"
	players[2] = "Doe"
	players[3] = "Victor"

	fmt.Println(players)

	for num, name := range players {
		fmt.Printf("Number: %d Player: %s\n", num, name)
	}

	fmt.Println(players[2])

	delete(players, 3)

	fmt.Println(players)
}
