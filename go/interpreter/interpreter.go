package interpreter

import (
	"fmt"
	"interpreter/expr"
	"interpreter/loxToken"
	"reflect"
)

type Interpreter struct{}

func (i *Interpreter) Something() any {
	return ""
}
func (i *Interpreter) VisitLiteralExpr(expr *expr.Literal) any {
	return expr.Value
}
func (i *Interpreter) VisitBinaryExpr(expr *expr.Binary) any {
	left, right := i.evaluate(expr.Left), i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case loxToken.BANG_EQUAL:
		return isEqual(left, right)
	case loxToken.EQUAL_EQUAL:
		return isEqual(left, right)
	case loxToken.PLUS:
		// Handle addition/concatenation
		return i.handleAddition(left, right)
	case loxToken.MINUS, loxToken.SLASH, loxToken.STAR:
		// Handle numeric operations
		return i.handleNumericOperation(expr.Operator.Type, left, right)
	default:
		return nil
	}
}

func (i *Interpreter) VisitGroupingExpr(expr *expr.Grouping) any {
	return i.evaluate(expr)
}
func (i *Interpreter) VisitUnaryExpr(expr *expr.Unary) any {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case loxToken.MINUS:
		if val, ok := toFloat64(right); ok {
			return -val
		}
		return nil
	case loxToken.BANG:
		return i.isTruthy(expr.Right)
	}
	return nil
}

func (i *Interpreter) evaluate(exp expr.IExpr[any]) any {
	return exp.Accept(i)
}

func (i *Interpreter) isTruthy(obj any) bool {
	if obj != nil {
		return false
	}
	if val, ok := obj.(bool); ok {
		return val
	}
	return true
}

func (i *Interpreter) handleAddition(left, right any) any {
	// Try numeric addition
	if leftNum, rightNum, ok := i.getNumericOperands(left, right); ok {
		return leftNum + rightNum
	}

	// Try string concatenation
	if res, ok := tryString(left, right); ok {
		return res
	}
	return nil
}

func tryString(a, b any) (string, bool) {
	left, okL := a.(string)
	right, okR := b.(string)
	if okL && okR {
		return left + right, true
	}
	if !(okL || okR) {
		return "", false

	}
	leftFloat, okLF := toFloat64(a)
	rightFloat, okRF := toFloat64(b)
	if okL && okLF {
		return fmt.Sprint(left, leftFloat), true
	} else if okL && okRF {
		return fmt.Sprint(left, rightFloat), true
	} else if okR && okLF {
		return fmt.Sprint(right, leftFloat), true

	} else if okR && okRF {
		return fmt.Sprint(right, rightFloat), true
	}
}

func (i *Interpreter) handleNumericOperation(operator loxToken.TokenType, left, right any) any {
	leftNum, rightNum, ok := i.getNumericOperands(left, right)
	if !ok {
		return nil
	}

	switch operator {
	case loxToken.GREATER:
		return leftNum > rightNum
	case loxToken.GREATER_EQUAL:
		return leftNum >= rightNum
	case loxToken.LESS:
		return leftNum < rightNum
	case loxToken.LESS_EQUAL:
		return leftNum <= rightNum
	case loxToken.MINUS:
		return leftNum - rightNum
	case loxToken.SLASH:
		return leftNum / rightNum
	case loxToken.STAR:
		return leftNum * rightNum
	default:
		return nil
	}
}

func isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}

func (i *Interpreter) getNumericOperands(left, right any) (float64, float64, bool) {
	leftNum, okLeft := toFloat64(left)
	rightNum, okRight := toFloat64(right)
	return leftNum, rightNum, okLeft && okRight
}

func toFloat64(v any) (float64, bool) {
	switch val := v.(type) {
	case int:
		return float64(val), true
	case int32:
		return float64(val), true
	case int64:
		return float64(val), true
	case float32:
		return float64(val), true
	case float64:
		return val, true
	default:
		return 0, false
	}
}
