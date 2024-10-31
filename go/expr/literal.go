package expr

type Literal struct {
	value any
}

func NewLiteral(val any) Literal {
	return Literal{value: val}
}

func (l *Literal) Accept(v IVisitor[string]) string {
	return v.VisitLiteralExpr(l)
}
