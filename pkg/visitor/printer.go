package visitor

import "fmt"

type printer struct {
}

func (p *printer) visitForA(a *a) {
	fmt.Println(a.name)
}

func (p *printer) visitForB(b *b) {
	fmt.Println(b.number)
}

func (p *printer) visitForC(someC *c) {
	fmt.Println(someC.flag)
}

//NewPrinter creates private implementation of Visitor interface
//Printer prints private data of structures
func NewPrinter() Visitor {
	return &printer{}
}
