package expr

type Literal struct {
	Value any
}

func (l *Literal) Accept(v IVisitor[any]) any {
	return v.VisitLiteralExpr(l)
}

func NewLiteral[T any](val T) IExpr[any] {
	return &Literal{Value: val}
}
