package lox

import (
	"interpreter/errorz"
	"interpreter/utils"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) Scanner {
	return Scanner{source: source, tokens: []Token{}, start: 0, current: 0, line: 1}
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanTokens()
	}
	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	r := s.advance()
	switch r {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		s.addToken(TokenType(utils.TernararyHelper(func() bool { return s.match('=') }, BANG_EQUAL, BANG)))
	case '=':
		s.addToken(TokenType(utils.TernararyHelper(func() bool { return s.match('=') }, EQUAL_EQUAL, EQUAL)))
	case '<':
		s.addToken(TokenType(utils.TernararyHelper(func() bool { return s.match('=') }, LESS_EQUAL, LESS)))
	case '>':
		s.addToken(TokenType(utils.TernararyHelper(func() bool { return s.match('=') }, GREATER_EQUAL, GREATER)))
	default:
		errorz.Error(s.line, "Unexpected character")
	}
}
func (s *Scanner) match(c rune) bool {
	if s.isAtEnd() {
		return false
	}
	if rune(s.source[s.current]) != c {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

func (s *Scanner) addToken(tokenType TokenType) {

}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal any) {
	input := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType: tokenType, literal: literal, lexeme: input, line: s.line})
}
