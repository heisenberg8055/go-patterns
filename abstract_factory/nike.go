package abstract_factory

// Concrete Factory
type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{Shoe{logo: "nike", size: 11}}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{Shirt: Shirt{logo: "nike", size: 41}}
}
