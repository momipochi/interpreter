package expr

type IVisitor[T any] interface {
	Something() T
	visitLiterlExpr(expr Literal) T
	visitBinaryExpr(expr Binary) T
}

type IExpr[T any] interface {
	Accept(visitor IVisitor[T]) string
}
