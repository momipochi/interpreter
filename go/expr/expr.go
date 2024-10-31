package expr

import "interpreter/lox"

type Expr struct {
	Left     *Expr
	Operator lox.Token
	Right    *Expr
}
