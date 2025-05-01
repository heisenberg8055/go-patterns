package visitor

type Shape interface {
	getType() string
	Accept(Visitor)
}
