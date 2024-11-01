package expr

import "interpreter/loxToken"

type Unary struct {
	Operator loxToken.Token
	Right    IExpr[any]
}

// Accept implements IExpr.
func (u *Unary) Accept(visitor IVisitor[any]) any {
	return visitor.VisitUnaryExpr(u)
}

func NewUnary(operator loxToken.Token, expression IExpr[any]) IExpr[any] {
	return &Unary{Operator: operator, Right: expression}
}
