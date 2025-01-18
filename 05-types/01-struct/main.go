package main

import "fmt"

// agrupador de dados
type Product struct {
	name     string
	price    float64
	discount float64
}

func (p Product) getPriceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	prod1 := Product{
		name:     "Mac Book",
		price:    5000,
		discount: 0.05,
	}

	fmt.Println(prod1.getPriceWithDiscount())
}
