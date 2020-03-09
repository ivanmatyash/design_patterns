package pizza

type Pepperoni struct{}

func (Pepperoni) Description() string {
	return "Pepperoni pizza"
}

func (Pepperoni) Cost() uint64 {
	return 18
}
