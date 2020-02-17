package flying

// implements Flyer interface
type NoFlyBehavior struct{}

func NewNoFlyBehavior() Flyer {
	return NoFlyBehavior{}
}

func (NoFlyBehavior) Fly() string {
	return "I can not fly!"
}
