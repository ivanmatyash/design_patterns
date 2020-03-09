package additives

import "github.com/ivanmatyash/design-patterns/decorator/pkg/pizza"

type Meat struct {
	pz pizza.Pizzer
}

func NewMeat(p pizza.Pizzer) pizza.Pizzer {
	return Meat{pz: p}
}

func (m Meat) Description() string {
	return m.pz.Description() + " Meat"
}

func (m Meat) Cost() uint64 {
	return m.pz.Cost() + 5
}
