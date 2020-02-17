package animal

import (
	"github.com/ivanmatyash/design-patterns/strategy/pkg/flying"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/swimming"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/walking"
)

type Duck struct {
	Animal
}

func NewDuck() Duck {
	return Duck{
		Animal{
			FlyBehavior:  flying.NewFlyWithWingsBehavior(2),
			SwimBehavior: swimming.NewSwimWithFinsBehavior(2),
			WalkBehavior: walking.NewWalkWithLegsBehavior(2),
		},
	}
}
