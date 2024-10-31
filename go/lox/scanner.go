package lox

import (
	"interpreter/errorz"
	"interpreter/utils"
	"strconv"
)

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	keywords map[string]TokenType
}

func NewScanner(source string) Scanner {
	keywords := map[string]TokenType{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
	return Scanner{source: source, tokens: []Token{}, start: 0, current: 0, line: 1, keywords: keywords}
}

func (s *Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
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
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, BANG_EQUAL, BANG))
	case '=':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, EQUAL_EQUAL, EQUAL))
	case '<':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, LESS_EQUAL, LESS))
	case '>':
		s.addToken(utils.TernararyHelper(func() bool { return s.match('=') }, GREATER_EQUAL, GREATER))
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
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
			s.addToken(OR)
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
		val = IDENTIFIER
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
	s.addTokenLiteral(NUMBER, val)
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
	s.addTokenLiteral(STRING, val)
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

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal any) {
	input := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType: tokenType, literal: literal, Lexeme: input, line: s.line})
}
