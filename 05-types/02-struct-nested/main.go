package main

import "fmt"

type Address struct {
	zipCode string
	street  string
	number  uint
}

type Person struct {
	name    string
	age     uint8
	gender  string
	address Address
}

func main() {
	person := Person{
		name:   "Antonio Victor",
		age:    22,
		gender: "male",
		address: Address{
			zipCode: "12312312",
			street:  "street",
			number:  3032,
		},
	}

	fmt.Println(person)
}
