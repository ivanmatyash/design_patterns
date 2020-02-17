package animal

import (
	"github.com/ivanmatyash/design-patterns/strategy/pkg/flying"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/swimming"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/walking"
)

type Shark struct {
	Animal
}

func NewShark() Shark {
	return Shark{
		Animal{
			FlyBehavior:  flying.NewNoFlyBehavior(),
			SwimBehavior: swimming.NewSwimWithFinsBehavior(2),
			WalkBehavior: walking.NewNoWalkBehavior(),
		},
	}
}
