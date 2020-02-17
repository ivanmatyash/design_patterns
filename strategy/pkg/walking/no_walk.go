package walking

type NoWalkBehavior struct{}

func NewNoWalkBehavior() Walker {
	return NoWalkBehavior{}
}

func (NoWalkBehavior) Walk() string {
	return "I can not walk!"
}
