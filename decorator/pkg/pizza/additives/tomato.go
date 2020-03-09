package additives

import "github.com/ivanmatyash/design-patterns/decorator/pkg/pizza"

type Tomato struct {
	pz pizza.Pizzer
}

func NewTomato(p pizza.Pizzer) pizza.Pizzer {
	return Tomato{pz: p}
}

func (t Tomato) Description() string {
	return t.pz.Description() + " Tomato"
}

func (t Tomato) Cost() uint64 {
	return t.pz.Cost() + 1
}
