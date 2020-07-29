package firstacceptortype

type visitor interface {
	VisitForFirstType(interface {
		GetStr() string
	})
	VisitForSecondType(interface {
		GetNumber() int
	})
}

//Acceptor describes how to delegate handling to Visitor
type Acceptor interface {
	Accept(visitor)
}

type firstAcceptorType struct {
	name string
}

func (f *firstAcceptorType) Accept(visitor visitor) {
	visitor.VisitForFirstType(f)
}

func (f *firstAcceptorType) GetStr() (str string) {
	str = f.name
	return
}

//NewFirstAcceptorType creates private implementation of Acceptor interface
func NewFirstAcceptorType(str string) Acceptor {
	return &firstAcceptorType{
		name: str,
	}
}
