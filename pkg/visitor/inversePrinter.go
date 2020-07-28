package visitor

import "fmt"

type inversePrinter struct {
}

func (p *inversePrinter) visitForA(*a) {
}

func (p *inversePrinter) visitForB(b *b) {
	fmt.Println((-1) * b.number)
}

func (p *inversePrinter) visitForC(someC *c) {
	fmt.Println(!someC.flag)
}

//NewInversePrinter creates private implementation of Visitor interface
//InversePrinter prints inverse values of private data
func NewInversePrinter() Visitor {
	return &inversePrinter{}
}
