package visitor

type b struct {
	number int64
}

func (b *b) Accept(v Visitor) {
	v.visitForB(b)
}

func NewB() Acceptor {
	return &b{
		number: 1,
	}
}
