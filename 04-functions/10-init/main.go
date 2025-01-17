package main

import "fmt"

// dentro de cada arquivo go você pode ter uma função init
func init() {
	fmt.Println("Init...")
}

// a func main deve ser única no programa
func main() {
	fmt.Println("Main...")
}
