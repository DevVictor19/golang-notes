package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name    string
	surname string
}

func (p Person) getFullName() string {
	return p.name + " " + p.surname
}

// para alterar valores do struct passamos ele por referência
func (p *Person) setFullName(fullname string) {
	parts := strings.Split(fullname, " ")
	p.name = parts[0]
	p.surname = parts[1]
}

func main() {
	p := Person{name: "Antonio", surname: "Oliveira"}

	fmt.Println(p.getFullName())

	p.setFullName("João Victor")

	fmt.Println(p.getFullName())
}
