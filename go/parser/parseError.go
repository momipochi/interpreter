package parser

import (
	"fmt"
	"interpreter/errorz"
	"interpreter/lox"
)

type ParseError struct {
	token   lox.Token
	message string
}

func NewParseError(token lox.Token, message string) *ParseError {
	return &ParseError{token: token, message: message}
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%s - %s", p.token, p.message)
}

func ParsingError(token lox.Token, message string) {
	if token.Type == lox.EOF {
		errorz.Report(token.Line, " at end", message)
	} else {
		errorz.Report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}
