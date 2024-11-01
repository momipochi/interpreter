package loxToken

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{Type: tokenType, Lexeme: lexeme, Literal: literal, Line: line}
}
func NewTokenNoLiteral(tokenType TokenType, literal any, line int) Token {
	return Token{Type: tokenType, Lexeme: string(tokenType), Literal: literal, Line: line}
}

func (t *Token) ToString() string {
	return fmt.Sprintf("Type: %s Lexem: %s Literal: %s", t.Type, t.Lexeme, t.Literal)
}
