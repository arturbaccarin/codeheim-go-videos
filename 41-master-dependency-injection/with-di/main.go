package main

import "fmt"

type Ingredient interface {
	Mix() string
}

type Flour struct{}

func (Flour) Mix() string {
	return "mixing flour"
}

type Sugar struct{}

func (Sugar) Mix() string {
	return "mixing sugar"
}

type Butter struct{}

func (Butter) Mix() string {
	return "mixing butter"
}

type Oven interface {
	Heat() string
}

type GasOven struct{}

func (GasOven) Heat() string {
	return "heating with gas oven"
}

type EletricOven struct{}

func (EletricOven) Heat() string {
	return "heating with eletric oven"
}

// Bakery depends on the Oven and Ingredients
type Bakery struct {
	oven        Oven
	ingredients []Ingredient
}

// Bake is a method that uses the oven and ingredients to bake a pastry
func (b *Bakery) Bake() {
	fmt.Println(b.oven.Heat())

	for _, ingredient := range b.ingredients {
		fmt.Println(ingredient.Mix())
	}

	fmt.Println("baking an awesome pastry")
}

func main() {
	gasOven := GasOven{}
	eletricOven := EletricOven{}
	flour := Flour{}
	sugar := Sugar{}
	butter := Butter{}

	bakery := Bakery{
		oven:        gasOven,
		ingredients: []Ingredient{flour, sugar, butter},
	}
	bakery.Bake()

	bakery = Bakery{
		oven:        eletricOven,
		ingredients: []Ingredient{flour, butter},
	}

	bakery.Bake()
}
