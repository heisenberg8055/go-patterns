package decorator

type TomatoToppings struct {
	Pizza IPizza
}

func (c *TomatoToppings) GetPrice() uint {
	return c.Pizza.GetPrice() + 7
}
