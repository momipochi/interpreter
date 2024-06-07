package expr

type Visitor[T any] interface {
	something() T
}

type Expr[T any] interface {
	Accept() T
}
