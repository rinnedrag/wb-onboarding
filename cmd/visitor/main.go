package main

import (
	"wb-onboarding/pkg/visitor/firstacceptortype"
	"wb-onboarding/pkg/visitor/printer"
	"wb-onboarding/pkg/visitor/secondacceptortype"
)

func main() {
	a := firstacceptortype.NewFirstAcceptorType("a")
	b := secondacceptortype.NewSecondAcceptorType(1)
	printerVisitor := printer.NewPrinter()
	a.Accept(printerVisitor)
	b.Accept(printerVisitor)
}
