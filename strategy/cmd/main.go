package main

import (
	"fmt"

	"github.com/ivanmatyash/design-patterns/strategy/pkg/animal"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/walking"
)

func main() {
	cat := animal.NewCat()
	fmt.Println(cat.PerformFly())
	fmt.Println(cat.PerformSwim())
	fmt.Println(cat.PerformWalk())

	cat.SetWalkBehavior(walking.NewWalkWithLegsBehavior(5))
	fmt.Println(cat.PerformWalk())

}
