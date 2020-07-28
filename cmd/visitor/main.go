package main

import "wb-onboarding/pkg/visitor"

func main() {
	a := visitor.NewA()
	b := visitor.NewB()
	c := visitor.NewC()
	container := []visitor.Acceptor{a, b, c}
	printer := visitor.NewPrinter()
	for _, el := range container {
		el.Accept(printer)
	}
	invPrinter := visitor.NewInversePrinter()
	for _, el := range container {
		el.Accept(invPrinter)
	}
}
