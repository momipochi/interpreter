package expr

type IVisitor[T any] interface {
	Something() T
	VisitLiteralExpr(expr *Literal) T
	VisitBinaryExpr(expr *Binary[T]) T
	VisitGroupingExpr(expr *Grouping[T]) T
	VisitUnaryExpr(expr *Unary[T]) T
}

type IExpr[T any] interface {
	Accept(visitor IVisitor[T]) T
}
