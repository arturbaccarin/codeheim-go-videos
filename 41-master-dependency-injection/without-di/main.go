package main

import "fmt"

// Bakery depends on the Oven and Ingredients
type Bakery struct {
	ovenType    string
	ingredients []string
}

// Bake is a method that uses the oven and ingredients to bake a pastry
func (b *Bakery) Bake() {
	switch b.ovenType {
	case "gas oven":
		fmt.Println("Heating with gas oven")
	case "eletrict oven":
		fmt.Println("Heating with eletrict oven")
	}

	for _, ingredient := range b.ingredients {
		switch ingredient {
		case "flour":
			fmt.Println("Adding flour")
		case "sugar":
			fmt.Println("Adding sugar")
		case "butter":
			fmt.Println("Adding butter")
		}
	}

	fmt.Println("baking an awesome pastry")
}

func main() {
	bakery := Bakery{
		ovenType:    "gas oven",
		ingredients: []string{"flour", "sugar", "butter"},
	}
	bakery.Bake()

	bakery = Bakery{
		ovenType:    "eletrict oven",
		ingredients: []string{"flour", "butter"},
	}
	bakery.Bake()
}
