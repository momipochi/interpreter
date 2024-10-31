package main

import (
	"fmt"
	astprinter "interpreter/astPrinter"
	"interpreter/expr"
	"interpreter/lox"
)

func main() {
	a := expr.NewUnary[any](lox.NewTokenNoLiteral(lox.MINUS, nil, 1), expr.NewLiteral(123))
	// var expression expr.IExpr[string] = expr.NewBinary[string](
	// 	,

	// 	lox.NewTokenNoLiteral(lox.STAR, nil, 1),

	// 	expr.NewGrouping[any](
	// 		expr.NewLiteral[any](45.67),
	// 	))

	printer := astprinter.NewPrinter()
	fmt.Println(printer.Print(expression))
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
