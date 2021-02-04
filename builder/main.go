package main

import (
	"fmt"
	"go-design-patterns/builder/pizza"
)

func main() {
	_, b := pizza.NewBuilder(pizza.Pepperoni)
	d := pizza.Director{}
	d.SetBuilder(b)
	p := d.GetPizza()
	fmt.Printf("%+v\n", p)
}
