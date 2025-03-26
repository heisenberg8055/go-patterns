package abstract_factory

// Abstract Product
type IShoe interface {
	setLogo(logo string)
	setSize(size uint8)
	getLogo() string
	getSize() uint8
}

type Shoe struct {
	logo string
	size uint8
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size uint8) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() uint8 {
	return s.size
}
