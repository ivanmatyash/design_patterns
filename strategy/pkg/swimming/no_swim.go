package swimming

type NoSwimBehavior struct{}

func NewNoSwimBehavior() Swimmer {
	return NoSwimBehavior{}
}

func (NoSwimBehavior) Swim() string {
	return "I can not swim!"
}
