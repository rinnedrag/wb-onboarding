package visitor

type a struct {
	name string
}

func (a *a) Accept(v Visitor) {
	v.visitForA(a)
}

func NewA() Acceptor {
	return &a{
		name: "a",
	}
}
