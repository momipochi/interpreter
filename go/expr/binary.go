package expr

import "interpreter/lox"

type Binary struct {
	Left     IExpr[any]
	Operator lox.Token
	Right    IExpr[any]
}

// Accept implements IExpr.
func (b *Binary) Accept(visitor IVisitor[any]) any {
	return visitor.VisitBinaryExpr(b)
}

func NewBinary(l IExpr[any], o lox.Token, r IExpr[any]) IExpr[any] {
	return &Binary{Left: l, Operator: o, Right: r}
}
