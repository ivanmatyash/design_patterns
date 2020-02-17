package flying

import "fmt"

// implements `Flyer` interface
type FlyWithWingsBehavior struct {
	amount int
}

func NewFlyWithWingsBehavior(wings int) Flyer {
	return FlyWithWingsBehavior{amount: wings}
}

func (f FlyWithWingsBehavior) Fly() string {
	return fmt.Sprintf("I am flying with %d wings!", f.amount)
}
