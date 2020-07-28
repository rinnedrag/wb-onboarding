package visitor

type Acceptor interface {
	Accept(Visitor)
}

type Visitor interface {
	visitForA(a *a)
	visitForB(b *b)
	visitForC(c *c)
}
