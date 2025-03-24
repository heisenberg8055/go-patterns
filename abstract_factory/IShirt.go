package abstract_factory

// Abstract Product
type IShirt interface {
	setLogo(logo string)
	setSize(size uint8)
	getLogo() string
	getSize() uint8
}

type Shirt struct {
	logo string
	size uint8
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size uint8) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() uint8 {
	return s.size
}
