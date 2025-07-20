package main

import "fmt"

type Order struct {
	Amount float64
}

type DiscountStrategy interface {
	Calculate(order Order) float64
}

type NormalDiscount struct{}

func (d NormalDiscount) Calculate(order Order) float64 {
	return order.Amount * 0.05
}

type BlackFridayDiscount struct{}

func (d BlackFridayDiscount) Calculate(order Order) float64 {
	return order.Amount * 0.30
}

type ClearanceDiscount struct{}

func (d ClearanceDiscount) Calculate(order Order) float64 {
	return order.Amount * 0.50
}

type DiscountContext struct {
	strategy DiscountStrategy
}

func (ctx DiscountContext) SetStrategy(strategy DiscountStrategy) {
	ctx.strategy = strategy
}

func (ctx DiscountContext) ApplyDiscount(order Order) float64 {
	if ctx.strategy == nil {
		return 0
	}
	return ctx.strategy.Calculate(order)
}

func main() {
	order := Order{Amount: 300}
	context := DiscountContext{}

	// Normal
	context.SetStrategy(NormalDiscount{})
	fmt.Printf("Normal discount: $%.2f\n", context.ApplyDiscount(order))

	// Black Friday
	context.SetStrategy(BlackFridayDiscount{})
	fmt.Printf("Black Friday discount: $%.2f\n", context.ApplyDiscount(order))

	// Clearance
	context.SetStrategy(ClearanceDiscount{})
	fmt.Printf("Clearance discount: $%.2f\n", context.ApplyDiscount(order))
}
