package astprinter

import (
	"interpreter/expr"
	"strings"
)

type AstPrinter struct{}

func NewPrinter() AstPrinter {
	return AstPrinter{}
}

func (ap *AstPrinter) Something() string {
	return ""
}
func (ap *AstPrinter) VisitLiteralExpr(expr *expr.Literal) string {
	return ""
}
func (ap *AstPrinter) VisitBinaryExpr(expr *expr.Binary[string]) string {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ap *AstPrinter) VisitGroupingExpr(expr *expr.Grouping[string]) string {
	return ap.parenthesize("group", expr.Expression)
}
func (ap *AstPrinter) VisitUnaryExpr(expr *expr.Unary[string]) string {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (ap *AstPrinter) parenthesize(name string, expr ...expr.IExpr[string]) string {
	sb := strings.Builder{}
	sb.WriteString("(")
	sb.WriteString(name)
	for _, val := range expr {
		sb.WriteString(val.Accept(ap))
	}
	sb.WriteString(")")
	return sb.String()
}

func (ap *AstPrinter) Print(expr expr.IExpr[string]) string {
	return expr.Accept(ap)
}
