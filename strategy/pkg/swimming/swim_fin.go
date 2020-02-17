package swimming

import "fmt"

type SwimWithFinsBehavior struct {
	amount int
}

func NewSwimWithFinsBehavior(fins int) Swimmer {
	return SwimWithFinsBehavior{amount: fins}
}

func (s SwimWithFinsBehavior) Swim() string {
	return fmt.Sprintf("I am swimmig with %d fins!", s.amount)
}
