package expr

type Literal struct{}

func NewLiteral() Literal {
	return Literal{}
}

func (l *Literal) Accept() {

}
func (b *Binary) visitBinary() string {
	return " "
}
