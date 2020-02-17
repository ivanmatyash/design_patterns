package walking

import "fmt"

// implements Walker interface
type WalkWithLegsBehavior struct {
	amount int
}

func NewWalkWithLegsBehavior(legs int) Walker {
	return WalkWithLegsBehavior{amount: legs}
}

func (w WalkWithLegsBehavior) Walk() string {
	return fmt.Sprintf("I am walking with %d legs!", w.amount)
}
