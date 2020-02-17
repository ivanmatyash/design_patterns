package animal

import (
	"github.com/ivanmatyash/design-patterns/strategy/pkg/flying"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/swimming"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/walking"
)

type Cat struct {
	Animal
}

func NewCat() Cat {
	return Cat{
		Animal{
			FlyBehavior:  flying.NewNoFlyBehavior(),
			SwimBehavior: swimming.NewNoSwimBehavior(),
			WalkBehavior: walking.NewWalkWithLegsBehavior(4),
		},
	}
}
