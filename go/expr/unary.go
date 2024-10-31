package expr

import "interpreter/lox"

type Unary struct {
	Operator lox.Token
	Right    IExpr[any]
}

// Accept implements IExpr.
func (u *Unary) Accept(visitor IVisitor[any]) any {
	return visitor.VisitUnaryExpr(u)
}

func NewUnary(op lox.Token, r IExpr[any]) IExpr[any] {
	return &Unary{Operator: op, Right: r}
}
