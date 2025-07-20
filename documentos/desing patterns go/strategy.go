package main

import (
	"fmt"
)

type Order struct {
	Amount float64
	Type   string // "normal", "black_friday", "clearance"
}

func CalculateDiscount(order Order) float64 {
	if order.Type == "normal" {
		return order.Amount * 0.05
	} else if order.Type == "black_friday" {
		return order.Amount * 0.30
	} else if order.Type == "clearance" {
		return order.Amount * 0.50
	}
	return 0
}

func main() {
	orders := []Order{
		{Amount: 100, Type: "normal"},
		{Amount: 200, Type: "black_friday"},
		{Amount: 300, Type: "clearance"},
	}

	for _, order := range orders {
		discount := CalculateDiscount(order)
		fmt.Printf("Order of type %s gets a discount of $%.2f\n", order.Type, discount)
	}
}
