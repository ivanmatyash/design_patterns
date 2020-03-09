package main

import (
	"fmt"

	"github.com/ivanmatyash/design-patterns/decorator/pkg/pizza"
	"github.com/ivanmatyash/design-patterns/decorator/pkg/pizza/additives"
)

func main() {
	margarita := pizza.Margarita{}
	mCheeseMeat := additives.NewMeat(additives.NewCheese(margarita))

	fmt.Println(mCheeseMeat.Description())
	fmt.Println(mCheeseMeat.Cost())

	p := additives.NewTomato(pizza.Pepperoni{})
	fmt.Println(p.Description())
	fmt.Println(p.Cost())
}
