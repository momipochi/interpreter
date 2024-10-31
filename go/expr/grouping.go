package expr

type Grouping[T any] struct {
	Expression IExpr[T]
}

func NewGrouping[T any](expr IExpr[T]) Grouping[T] {
	return Grouping[T]{Expression: expr}
}
func (g *Grouping[any]) Accept(v IVisitor[any]) any {
	return v.VisitGroupingExpr(g)
}
