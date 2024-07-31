package main

import "fmt"

// Product
type Pizza struct {
	Size      string
	Crust     string
	Cheese    bool
	Pepperoni bool
	Mushrooms bool
}

type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	Build() Pizza
}

// Implements PizzaBuilder interface
type ConcretePizzaBuilder struct {
	pizza Pizza
}

func (b *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

func (b *ConcretePizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.Cheese = true
	return b
}

func (b *ConcretePizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}

func (b *ConcretePizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.Mushrooms = true
	return b
}

func (b *ConcretePizzaBuilder) Build() Pizza {
	return b.pizza
}

// Define commons pizzas configurations
type PizzaDirector struct{}

func (p *PizzaDirector) CreateMargherita(b PizzaBuilder) Pizza {
	b.SetSize("Small").SetCrust("Thin").AddCheese()
	return b.Build()
}

// Other predefines pizzas can be added...

func main() {
	builder := &ConcretePizzaBuilder{}
	director := PizzaDirector{}

	// Predefined Pizza
	margherita := director.CreateMargherita(builder)
	fmt.Println("Margherita: ", margherita)

	// Custom Pizza
	customPizza := builder.SetSize("Large").AddPepperoni().AddMushrooms().Build()
	fmt.Println("Custom Pizza: ", customPizza)
}
