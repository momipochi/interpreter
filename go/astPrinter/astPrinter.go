package astprinter

import (
	"interpreter/expr"
	"strings"
)

type AstPrinter struct{}

func (ap *AstPrinter) Something() string {
	return ""
}

func (ap *AstPrinter) parenthesize(name string, expr []expr.IExpr[string]) string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(name)
	for _, val := range expr {
		sb.WriteString(val.Accept(ap))
	}
	sb.WriteString(")")
	return sb.String()
}
