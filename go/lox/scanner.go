package lox

import (
	"interpreter/errorz"
	"interpreter/loxToken"
	"interpreter/utils"
	"strconv"
)

type Scanner struct {
	source   string
	tokens   []loxToken.Token
	start    int
	current  int
	line     int
	keywords map[string]loxToken.TokenType
}

func NewScanner(source string) Scanner {
	keywords := map[string]loxToken.TokenType{
		"and":    loxToken.AND,
		"class":  loxToken.CLASS,
		"else":   loxToken.ELSE,
		"false":  loxToken.FALSE,
		"for":    loxToken.FOR,
		"fun":    loxToken.FUN,
		"if":     loxToken.IF,
		"nil":    loxToken.NIL,
		"or":     loxToken.OR,
		"print":  loxToken.PRINT,
		"return": loxToken.RETURN,
		"super":  loxToken.SUPER,
		"this":   loxToken.THIS,
		"true":   loxToken.TRUE,
		"var":    loxToken.VAR,
		"while":  loxToken.WHILE,
	}
	return Scanner{source: source, tokens: []loxToken.Token{}, start: 0, current: 0, line: 1, keywords: keywords}
}

func (s *Scanner) scanTokens() []loxToken.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, loxToken.NewToken(loxToken.EOF, "", nil, s.line))
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	r := s.advance()
	switch r {
	case '(':
		s.addToken(loxToken.LEFT_PAREN)
	case ')':
		s.addToken(loxToken.RIGHT_PAREN)
	case '{':
		s.addToken(loxToken.LEFT_BRACE)
	case '}':
		s.addToken(loxToken.RIGHT_BRACE)
	case ',':
		s.addToken(loxToken.COMMA)
	case '.':
		s.addToken(loxToken.DOT)
	case '-':
		s.addToken(loxToken.MINUS)
	case '+':
		s.addToken(loxToken.PLUS)
	case ';':
		s.addToken(loxToken.SEMICOLON)
	case '*':
		s.addToken(loxToken.STAR)
	case '!':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, loxToken.BANG_EQUAL, loxToken.BANG))
	case '=':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, loxToken.EQUAL_EQUAL, loxToken.EQUAL))
	case '<':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, loxToken.LESS_EQUAL, loxToken.LESS))
	case '>':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, loxToken.GREATER_EQUAL, loxToken.GREATER))
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(loxToken.SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.string()
	case 'o':
		if s.peek() == 'r' {
			s.addToken(loxToken.OR)
		}
	default:
		if utils.IsDigit(r) {
			s.number()
		} else if utils.IsAlpha(r) {
			s.identifier()
		} else {
			errorz.Error(s.line, "Unexpected character.")
		}
	}
}

func (s *Scanner) identifier() {
	for utils.IsAlphaNumeric(s.peek()) {
		s.advance()
	}
	val, ok := s.keywords[s.source[s.start:s.current]]
	if ok {
		val = loxToken.IDENTIFIER
	}
	s.addToken(val)
}

func (s *Scanner) number() {
	for utils.IsDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && utils.IsDigit(s.peekNext()) {
		s.advance()
		for utils.IsDigit(s.peek()) {
			s.advance()
		}
	}
	val, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	errorz.CheckError(err)
	s.addTokenLiteral(loxToken.NUMBER, val)
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '0'
	}
	return rune(s.source[s.current+1])
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		errorz.Error(s.line, "Unterminated string.")
		return
	}
	s.advance()
	val := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(loxToken.STRING, val)
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '0'
	}
	return rune(s.source[s.current])
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

func (s *Scanner) addToken(tokenType loxToken.TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType loxToken.TokenType, literal any) {
	input := s.source[s.start:s.current]
	s.tokens = append(s.tokens, loxToken.Token{Type: tokenType, Literal: literal, Lexeme: input, Line: s.line})
}
