package main

import "fmt"

// Component interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Concrete Component
type SimpleCoffee struct{}

func (*SimpleCoffee) Cost() float64 {
	return 1.0
}

func (*SimpleCoffee) Description() string {
	return "simple coffee"
}

// Decorator 1: Milk
type Milk struct {
	coffee Coffee
}

func (m *Milk) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.coffee.Description() + ", milk"
}

// Decorator 2: Caramel
type Caramel struct {
	coffee Coffee
}

func (c *Caramel) Cost() float64 {
	return c.coffee.Cost() + 1.0
}

func (c *Caramel) Description() string {
	return c.coffee.Description() + ", caramel"
}

func main() {
	// Create a simple coffee
	coffee := &SimpleCoffee{}
	fmt.Println("Coffee: ", coffee.Description(), " - Cost: ", coffee.Cost())

	// Decorator it with Milk
	coffeeWithMilk := &Milk{coffee: coffee}
	fmt.Println("Coffee with Milk: ", coffeeWithMilk.Description(), " - Cost: ", coffeeWithMilk.Cost())

	// Decorator it with Caramel
	coffeeWithCaramel := &Caramel{coffee: coffee}
	fmt.Println("Coffee with Caramel: ", coffeeWithCaramel.Description(), " - Cost: ", coffeeWithCaramel.Cost())

	// Decorator it with Milk and Caramel
	coffeeWithMilkAndCaramel := &Milk{coffee: &Caramel{coffee: coffee}}
	fmt.Println("Coffee with Milk and Caramel: ", coffeeWithMilkAndCaramel.Description(), " - Cost: ", coffeeWithMilkAndCaramel.Cost())
}
