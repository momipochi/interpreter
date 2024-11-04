package parser

import (
	"fmt"
	"interpreter/expr"
	"interpreter/loxToken"
)

type Parser struct {
	tokens  []loxToken.Token
	current int
}

func NewParser(t []loxToken.Token) *Parser {
	return &Parser{tokens: t}
}

func (p *Parser) expression() expr.IExpr[any] {
	res := p.equality()
	return res
}
func (p *Parser) equality() expr.IExpr[any] {
	expresssion := p.comparison()

	for p.match(loxToken.BANG_EQUAL, loxToken.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expresssion = expr.NewBinary(expresssion, *operator, right)

	}
	return expresssion
}

func (p *Parser) match(types ...loxToken.TokenType) bool {

	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(typ loxToken.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == typ
}

func (p *Parser) advance() *loxToken.Token {

	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}
func (p *Parser) isAtEnd() bool { return p.peek().Type == loxToken.EOF }
func (p *Parser) peek() *loxToken.Token {

	return &p.tokens[p.current]
}
func (p *Parser) previous() *loxToken.Token {
	return &p.tokens[p.current-1]
}
func (p *Parser) comparison() expr.IExpr[any] {
	expression := p.term()

	for p.match(loxToken.GREATER, loxToken.GREATER_EQUAL, loxToken.LESS, loxToken.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expression = expr.NewBinary(expression, *operator, right)
	}
	return expression
}

func (p *Parser) term() expr.IExpr[any] {
	expression := p.factor()

	for p.match(loxToken.MINUS, loxToken.PLUS) {
		operator := p.previous()
		right := p.factor()
		expression = expr.NewBinary(expression, *operator, right)
	}
	return expression
}
func (p *Parser) factor() expr.IExpr[any] {
	expression := p.unary()
	for p.match(loxToken.SLASH, loxToken.STAR) {
		operator := p.previous()
		right := p.unary()
		expression = expr.NewBinary(expression, *operator, right)
	}
	return expression
}
func (p *Parser) unary() expr.IExpr[any] {
	if p.match(loxToken.BANG, loxToken.MINUS) {
		operator := p.previous()
		right := p.unary()
		return expr.NewUnary(*operator, right)
	}
	return p.primary()
}
func (p *Parser) primary() expr.IExpr[any] {
	if p.match(loxToken.TRUE) {
		return expr.NewLiteral(true)
	}
	if p.match(loxToken.FALSE) {
		return expr.NewLiteral(false)
	}
	if p.match(loxToken.NIL) {
		return expr.NewLiteral("nil")
	}
	if p.match(loxToken.NUMBER, loxToken.STRING) {
		return expr.NewLiteral(p.previous().Literal)
	}
	if p.match(loxToken.LEFT_PAREN) {
		expression := p.expression()
		p.consume(loxToken.RIGHT_PAREN, "Expect ')' after expression")
		return expr.NewGrouping(expression)
	}
	return nil
}
func (p *Parser) consume(typ loxToken.TokenType, message string) *loxToken.Token {
	if p.check(typ) {
		return p.advance()
	}

	panic(p.error(*p.peek(), message))
}

func (p *Parser) error(token loxToken.Token, message string) *ParseError {
	ParsingError(token, message)
	return NewParseError(token, message)
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == loxToken.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case loxToken.CLASS:
		case loxToken.FUN:
		case loxToken.VAR:
		case loxToken.FOR:
		case loxToken.IF:
		case loxToken.WHILE:
		case loxToken.PRINT:
		case loxToken.RETURN:
			return
		}
		p.advance()
	}
}

func (p *Parser) Parse() (expression expr.IExpr[any], err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Paniced! Recovering...")
			if myErr, ok := r.(ParseError); ok {
				res := myErr.Error()
				err = fmt.Errorf("[parse error] error code %s", res)
			} else if myErr, ok := r.(RunTimeError); ok {
				res := myErr.Error()
				err = fmt.Errorf("[runtime error] error code %s", res)
			} else {
				err = fmt.Errorf("unknown panic: %v", r)
			}
		}
	}()

	res := p.expression()
	return res, nil
}

func (p *Parser) PrintContent() {
	fmt.Printf("Printing content \n %v", p.tokens)
}
