package additives

import "github.com/ivanmatyash/design-patterns/decorator/pkg/pizza"

type Cheese struct {
	pz pizza.Pizzer
}

func NewCheese(p pizza.Pizzer) pizza.Pizzer {
	return Cheese{pz: p}
}

func (c Cheese) Description() string {
	return c.pz.Description() + " Cheese"
}

func (c Cheese) Cost() uint64 {
	return c.pz.Cost() + 2
}
