package expr

import (
	"strings"
)

type Literal struct{}

func NewLiteral() Literal {
	return Literal{}
}

func (l *Literal) Accept() {

}
func (b *Binary) visitBinary() string {
	return parenthesize()
}

func parenthesize(name string, expr []Expr[string]) string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(name)
	for _, val := range expr {
		sb.WriteString(val.Accept())
	}
	sb.WriteString(")")
	return sb.String()
}
