package expr

import "interpreter/lox"

type Unary[T any] struct {
	Operator lox.Token
	Right    IExpr[T]
}

func NewUnary[T any](op lox.Token, r IExpr[T]) Unary[T] {
	return Unary[T]{Operator: op, Right: r}
}
func (u *Unary[any]) Accept(v IVisitor[any]) any {
	return v.VisitUnaryExpr(u)
}
