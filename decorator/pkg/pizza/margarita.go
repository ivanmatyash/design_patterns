package pizza

type Margarita struct{}

func (Margarita) Description() string {
	return "Margarita pizza"
}

func (Margarita) Cost() uint64 {
	return 15
}
