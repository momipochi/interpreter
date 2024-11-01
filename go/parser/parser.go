package parser

import (
	"interpreter/expr"
	"interpreter/lox"
)

type Parser struct {
	tokens  []lox.Token
	current int
}

func NewParser(t []lox.Token) *Parser {
	return &Parser{tokens: t}
}

func (p *Parser) expression() (expr.IExpr[any], error) {
	return p.equality()
}
func (p *Parser) equality() (expr.IExpr[any], error) {
	expresssion, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(lox.BANG_EQUAL, lox.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expresssion = expr.NewBinary(expresssion, *operator, right)
	}
	return expresssion, nil
}

func (p *Parser) match(types ...lox.TokenType) bool {
	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(typ lox.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == typ
}

func (p *Parser) advance() *lox.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}
func (p *Parser) isAtEnd() bool        { return p.peek().Type == lox.EOF }
func (p *Parser) peek() *lox.Token     { return &p.tokens[p.current] }
func (p *Parser) previous() *lox.Token { return &p.tokens[p.current-1] }
func (p *Parser) comparison() (expr.IExpr[any], error) {
	expression, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(lox.GREATER, lox.GREATER_EQUAL, lox.LESS, lox.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expression = expr.NewBinary(expression, *operator, right)
	}
	return expression, nil
}

func (p *Parser) term() (expr.IExpr[any], error) {
	expression, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(lox.MINUS, lox.PLUS) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expression = expr.NewBinary(expression, *operator, right)
	}
	return expression, nil
}
func (p *Parser) factor() (expr.IExpr[any], error) {
	expression, _ := p.unary()
	for p.match(lox.SLASH, lox.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err == nil {
			expression = expr.NewBinary(expression, *operator, right)

		}
		return nil, err
	}
	return expression, nil
}
func (p *Parser) unary() (expr.IExpr[any], error) {
	if p.match(lox.BANG, lox.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err == nil {
			return expr.NewUnary(*operator, right), nil
		}
		return nil, err

	}
	return p.primary()
}
func (p *Parser) primary() (expr.IExpr[any], error) {
	if p.match(lox.TRUE) {
		return expr.NewLiteral(true), nil
	}
	if p.match(lox.FALSE) {
		return expr.NewLiteral(false), nil
	}
	if p.match(lox.NIL) {
		return expr.NewLiteral("nil"), nil
	}
	if p.match(lox.NUMBER, lox.STRING) {
		return expr.NewLiteral(p.previous().Literal), nil
	}
	if p.match(lox.LEFT_PAREN) {
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume(lox.RIGHT_PAREN, "Expect ')' after expression")
		return expr.NewGrouping(expression), nil
	}
	return nil, p.error(*p.peek(), "Expect expresion.")
}
func (p *Parser) consume(typ lox.TokenType, message string) (*lox.Token, error) {
	if p.check(typ) {
		return p.advance(), nil
	}
	return nil, p.error(*p.peek(), message)
}

func (p *Parser) error(token lox.Token, message string) *ParseError {
	ParsingError(token, message)
	return NewParseError(token, message)
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == lox.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case lox.CLASS:
		case lox.FUN:
		case lox.VAR:
		case lox.FOR:
		case lox.IF:
		case lox.WHILE:
		case lox.PRINT:
		case lox.RETURN:
			return
		}
		p.advance()
	}
}

func (p *Parser) parse() (expr.IExpr[any], error) {
	val, err := p.expression()
	if err != nil {
		return nil, err
	}
	return val, nil
}
