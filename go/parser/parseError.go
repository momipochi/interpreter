package parser

import (
	"fmt"
	"interpreter/errorz"
	"interpreter/loxToken"
)

type ParseError struct {
	token   loxToken.Token
	message string
}

func NewParseError(token loxToken.Token, message string) *ParseError {
	return &ParseError{token: token, message: message}
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%v - %s", p.token, p.message)
}

func ParsingError(token loxToken.Token, message string) {
	if token.Type == loxToken.EOF {
		errorz.Report(token.Line, " at end", message)
	} else {
		errorz.Report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}
