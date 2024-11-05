package parser

import (
	"interpreter/errorz"
	"interpreter/loxToken"
)

type RunTimeError struct {
	token   loxToken.Token
	message string
}

func (p *RunTimeError) Error() string {
	return errorz.ReportToString(p.token.Line, "", p.message)
}

func NewRunTimeError(token loxToken.Token, message string) *RunTimeError {
	return &RunTimeError{token: token, message: message}
}
