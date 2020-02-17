package animal

import (
	"github.com/ivanmatyash/design-patterns/strategy/pkg/flying"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/swimming"
	"github.com/ivanmatyash/design-patterns/strategy/pkg/walking"
)

type Animal struct {
	FlyBehavior  flying.Flyer
	SwimBehavior swimming.Swimmer
	WalkBehavior walking.Walker
}

func (a *Animal) PerformFly() string {
	return a.FlyBehavior.Fly()
}

func (a *Animal) PerformSwim() string {
	return a.SwimBehavior.Swim()
}

func (a *Animal) PerformWalk() string {
	return a.WalkBehavior.Walk()
}

func (a *Animal) SetFlyBehavior(f flying.Flyer) {
	a.FlyBehavior = f
}

func (a *Animal) SetSwimBehavior(s swimming.Swimmer) {
	a.SwimBehavior = s
}

func (a *Animal) SetWalkBehavior(w walking.Walker) {
	a.WalkBehavior = w
}
