package expr

import (
	"interpreter/loxToken"
)

type Binary struct {
	Left     IExpr[any]
	Operator loxToken.Token
	Right    IExpr[any]
}

// Accept implements IExpr.
func (b *Binary) Accept(visitor IVisitor[any]) any {
	return visitor.VisitBinaryExpr(b)
}

func NewBinary(leftExpression IExpr[any], operator loxToken.Token, rightExpression IExpr[any]) IExpr[any] {
	return &Binary{Left: leftExpression, Operator: operator, Right: rightExpression}
}
