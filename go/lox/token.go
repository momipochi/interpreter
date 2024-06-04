package lox

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{tokenType: tokenType, lexeme: lexeme, literal: literal, line: line}
}

func (t *Token) toString() string {
	return fmt.Sprintf("Type: %s Lexem: %s Literal: %s", t.tokenType, t.lexeme, t.literal)
}
