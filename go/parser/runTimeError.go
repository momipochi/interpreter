package parser

import (
	"fmt"
	"interpreter/loxToken"
)

type RunTimeError struct {
	token   loxToken.Token
	message string
}

func (p *RunTimeError) Error() string {
	return fmt.Sprintf("%v \n[line %d] %s", p.token, p.token.Line, p.message)
}

func NewRunTimeError(token loxToken.Token, message string) *RunTimeError {
	return &RunTimeError{token: token, message: message}
}
