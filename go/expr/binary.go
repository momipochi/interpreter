package expr

import "interpreter/lox"

type Binary[T any] struct {
	Left     IExpr[T]
	Operator lox.Token
	Right    IExpr[T]
}

func NewBinary[T any]() Binary[T] {
	return Binary[T]{}
}

func (b *Binary[any]) Accept(v IVisitor[any]) any {
	return v.VisitBinaryExpr(b)
}
