package main

import "fmt"

type Person interface {
	sayHello()
}

// composition
type Student interface {
	Person
	study()
}

type Antonio struct{}

func (a Antonio) sayHello() {
	fmt.Println("Hello")
}

func (a Antonio) study() {
	fmt.Println("Studying...")
}

func main() {
	var student Student = Antonio{}
	student.sayHello()
	student.study()
}
