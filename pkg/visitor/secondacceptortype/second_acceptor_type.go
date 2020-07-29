package secondacceptortype

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

type secondAcceptorType struct {
	number int
}

func (s *secondAcceptorType) Accept(visitor visitor) {
	visitor.VisitForSecondType(s)
}

func (s *secondAcceptorType) GetNumber() (number int) {
	number = s.number
	return
}

//NewSecondAcceptorType creates private implementation of Acceptor interface
func NewSecondAcceptorType(number int) Acceptor {
	return &secondAcceptorType{
		number: number,
	}
}
