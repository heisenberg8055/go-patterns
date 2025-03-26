package abstract_factory

// Concrete Factory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe{logo: "adidas", size: 12},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt{logo: "adidas", size: 42},
	}
}
