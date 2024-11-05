package interpreter

import (
	"fmt"
	"interpreter/expr"
	"interpreter/loxToken"
	"interpreter/parser"
	"interpreter/utils"
	"reflect"
)

type Interpreter struct {
	hadError        *bool
	hadRuntimeError *bool
}

func NewInterpreter(hadErr, hadRunTimeErr *bool) *Interpreter {
	return &Interpreter{hadError: hadErr, hadRuntimeError: hadRunTimeErr}
}

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
		return i.handleAddition(left, right, expr.Operator)
	case loxToken.MINUS, loxToken.SLASH, loxToken.STAR, loxToken.GREATER, loxToken.GREATER_EQUAL, loxToken.LESS, loxToken.LESS_EQUAL:

		// Handle numeric operations
		return i.handleNumericOperation(expr.Operator, expr.Operator.Type, left, right)
	default:
		panic_NAN_NASTR(expr.Operator)
		return nil
	}
}

func (i *Interpreter) VisitGroupingExpr(expr *expr.Grouping) any {
	return i.evaluate(expr.Expression)
}
func (i *Interpreter) VisitUnaryExpr(expr *expr.Unary) any {
	right := i.evaluate(expr.Right)
	switch expr.Operator.Type {
	case loxToken.BANG:
		return i.isTruthy(expr.Right)
	case loxToken.MINUS:
		if val, ok := toFloat64(right); ok {
			return -val
		}
	}
	panic_NAN_NASTR(expr.Operator)
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
	return false
}

func (i *Interpreter) handleAddition(left, right any, token loxToken.Token) any {
	// Try numeric addition
	if leftNum, rightNum, ok := i.getNumericOperands(left, right); ok {
		return leftNum + rightNum
	}

	// Try string concatenation
	if res, ok := tryString(left, right); ok {
		return res
	}
	panic_NAN_NASTR(token)
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
	} else {
		return fmt.Sprint(right, rightFloat), true
	}
}

func (i *Interpreter) handleNumericOperation(token loxToken.Token, operator loxToken.TokenType, left, right any) any {
	leftNum, rightNum, ok := i.getNumericOperands(left, right)
	if !ok {
		panic_NAN_NASTR(token)
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
		if leftNum == 0 || rightNum == 0 {
			panic_DIVISION_BY_ZERO(token)
		}
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

func panic_NAN_NASTR(token loxToken.Token) {
	panic(parser.NewRunTimeError(token, "Operand not a number nor a string"))
}
func panic_DIVISION_BY_ZERO(token loxToken.Token) {
	panic(parser.NewRunTimeError(token, "Cannot divide by zero"))
}

func (i *Interpreter) Interpret(exp *expr.IExpr[any]) {
	val := i.evaluate(*exp)
	fmt.Println(stringify(val))

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(parser.RunTimeError); ok {
				fmt.Println(err.Error())
				*i.hadRuntimeError = true
			} else {
				panic(r)
			}
		}
	}()
}

func stringify(val any) any {
	if val == nil {
		return "nil"
	}
	if v, ok := toFloat64(val); ok {
		text := fmt.Sprint(v)
		if utils.StrEndsWith(text, ".0") {
			text = text[:len(text)-2]
		}
		return text
	}
	return fmt.Sprintln(val)
}
