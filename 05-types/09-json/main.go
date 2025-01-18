package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := Product{1, "Notebook", 1890, []string{"Discount", "For home", "Electronics"}}
	p1Json, _ := json.Marshal(p1)

	p1JsonStr := string(p1Json)

	fmt.Println(p1JsonStr)

	var p2 Product
	json.Unmarshal([]byte(p1JsonStr), &p2)

	fmt.Println(p2)
}
