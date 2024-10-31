package astprinter

import (
	"fmt"
	"interpreter/expr"
	"strings"
)

type AstPrinter struct{}

func NewPrinter() AstPrinter {
	return AstPrinter{}
}

func (ap *AstPrinter) Something() any {
	return ""
}
func (ap *AstPrinter) VisitLiteralExpr(expr *expr.Literal) any {
	if expr != nil {
		return expr.Value
	}
	return "nil"
}
func (ap *AstPrinter) VisitBinaryExpr(expr *expr.Binary) any {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ap *AstPrinter) VisitGroupingExpr(expr *expr.Grouping) any {
	return ap.parenthesize("group", expr.Expression)
}
func (ap *AstPrinter) VisitUnaryExpr(expr *expr.Unary) any {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (ap *AstPrinter) parenthesize(name string, expr ...expr.IExpr[any]) string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(name)
	for _, val := range expr {
		sb.WriteString(" ")
		sb.WriteString(fmt.Sprintf("%v", val.Accept(ap)))
	}
	sb.WriteString(")")
	return sb.String()
}

func (ap *AstPrinter) Print(expr *expr.IExpr[any]) any {
	return (*expr).Accept(ap)
}
