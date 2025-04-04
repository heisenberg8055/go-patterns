package decorator

type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() uint {
	return c.Pizza.GetPrice() + 8
}
