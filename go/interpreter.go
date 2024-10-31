package main

import (
	"fmt"
	astprinter "interpreter/astPrinter"
	"interpreter/expr"
	"interpreter/lox"
)

func main() {
	a := expr.NewUnary(lox.NewTokenNoLiteral(lox.MINUS, nil, 1), expr.NewLiteral(123))
	b := lox.NewTokenNoLiteral(lox.STAR, nil, 1)
	c := expr.NewGrouping(
		expr.NewLiteral(45.67),
	)
	expression := expr.NewBinary(a, b, c)
	// var expression expr.IExpr[string] = expr.NewBinary[string](
	// 	,

	// 	lox.NewTokenNoLiteral(lox.STAR, nil, 1),

	printer := astprinter.NewPrinter()
	fmt.Println(printer.Print(&expression))
	// for _, val := range os.Args {
	// 	fmt.Println(val)
	// }
	// l := lox.NewLox()
	// if len(os.Args) > 1 {
	// 	l.RunFile(os.Args[1])
	// } else {
	// 	l.RunPrompt()
	// }
}
