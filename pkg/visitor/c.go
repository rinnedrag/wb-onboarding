package visitor

type c struct {
	flag bool
}

func (c *c) Accept(v Visitor) {
	v.visitForC(c)
}

func NewC() Acceptor {
	return &c{
		flag: true,
	}
}
