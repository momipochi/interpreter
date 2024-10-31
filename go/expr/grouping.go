package expr

type Grouping struct {
	Expression IExpr[any]
}

// Accept implements IExpr.
func (g *Grouping) Accept(visitor IVisitor[any]) any {
	return visitor.VisitGroupingExpr(g)
}

func NewGrouping(expr IExpr[any]) IExpr[any] {
	return &Grouping{Expression: expr}
}
