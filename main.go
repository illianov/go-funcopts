package main

import (
	"fmt"
)

const (
	Pepperoni = "pepperoni"
	Mushrooms = "mushrooms"
	Anchovies = "anchovies"
	Pineapple = "pineapple"
)

var availToppings = map[string]bool{
	Pepperoni: true,
	Mushrooms: false,
	Anchovies: true,
	Pineapple: false,
}

func main()  {
	pizzaOrder := TakePizzaOrder()
	pizzaOrder.Print()
}

// TakePizzaOrder
func TakePizzaOrder() *PizzaOrder {
	pizzaOrder := NewPizzaOrder(
		WithTopping(Pepperoni),
		WithTopping(Mushrooms),
		WithTopping(Anchovies),
		WithTopping(Pineapple),
		)
	return pizzaOrder
}

type AddToppingFunc func(*PizzaOrder)

// NewPizzaOrder
func NewPizzaOrder(addToppingFuncs ...AddToppingFunc) *PizzaOrder {
	pizzaOrder := &PizzaOrder{
		toppingsAdded: make([]string, 0, 5),
		toppingsUnavail: make([]string, 0, 5),
	}

	for _, addToppingFunc := range addToppingFuncs {
		addToppingFunc(pizzaOrder)
	}
	return pizzaOrder
}

// WithTopping
func WithTopping(topping string) AddToppingFunc {
	return func(pizzaOrder *PizzaOrder) {
		pizzaOrder.AddTopping(topping)
	}
}

// PizzaOrder holds the pizza toppings
type PizzaOrder struct {
	toppingsAdded []string
	toppingsUnavail []string
}

// AddTopping
func (p *PizzaOrder) AddTopping(topping string) {
	avail := availToppings[topping]
	if avail == false {
		p.toppingsUnavail = append(p.toppingsUnavail, topping)
	} else {
		p.toppingsAdded = append(p.toppingsAdded, topping)
	}
}

func (p *PizzaOrder) Print() {
	fmt.Printf("Added toppings: %s\n", p.toppingsAdded)
	fmt.Printf("Unavailable toppings: %s\n", p.toppingsUnavail)
}


