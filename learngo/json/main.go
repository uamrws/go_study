package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID         string  `json:"id"`
	Name       string  `json:"name,omitempty"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

var O = Order{
	ID:         "1234",
	Quantity:   3,
	TotalPrice: 30.00,
}

func main() {
	s, err := json.Marshal(O)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
