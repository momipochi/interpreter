package lox

import "fmt"

type Token struct {
	tokenType TokenType
	Lexeme    string
	literal   any
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{tokenType: tokenType, Lexeme: lexeme, literal: literal, line: line}
}
func NewTokenNoLiteral(tokenType TokenType, literal any, line int) Token {
	return Token{tokenType: tokenType, Lexeme: string(tokenType), literal: literal, line: line}
}

func (t *Token) toString() string {
	return fmt.Sprintf("Type: %s Lexem: %s Literal: %s", t.tokenType, t.Lexeme, t.literal)
}
