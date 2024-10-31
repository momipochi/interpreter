package expr

type IVisitor[T any] interface {
	Something() T
	VisitLiteralExpr(expr *Literal) T
	VisitBinaryExpr(expr *Binary) T
	VisitGroupingExpr(expr *Grouping) T
	VisitUnaryExpr(expr *Unary) T
}

type IExpr[T any] interface {
	Accept(visitor IVisitor[T]) T
}
