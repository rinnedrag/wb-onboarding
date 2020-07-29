package printer

import "fmt"

//Visitor describes how to handle a component of certain type
type Visitor interface {
	VisitForFirstType(interface {
		GetStr() string
	})
	VisitForSecondType(interface {
		GetNumber() int
	})
}

type printer struct {
}

func (p *printer) VisitForFirstType(acceptor interface {
	GetStr() string
}) {
	fmt.Println(fmt.Sprintf("string: %s", acceptor.GetStr()))
}

func (p *printer) VisitForSecondType(acceptor interface {
	GetNumber() int
}) {
	fmt.Println(fmt.Sprintf("number: %d", acceptor.GetNumber()))
}

//NewPrinter creates private implementation of Visitor interface
//Printer prints private data of acceptors
func NewPrinter() Visitor {
	return &printer{}
}
