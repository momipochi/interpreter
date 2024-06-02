package lox

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
		break
	case ')':
		s.addToken(RIGHT_PAREN)
		break
	case '{':
		s.addToken(LEFT_BRACE)
		break
	case '}':
		s.addToken(RIGHT_BRACE)
		break
	case ',':
		s.addToken(COMMA)
		break
	case '.':
		s.addToken(DOT)
		break
	case '-':
		s.addToken(MINUS)
		break
	case '+':
		s.addToken(PLUS)
		break
	case ';':
		s.addToken(SEMICOLON)
		break
	case '*':
		s.addToken(STAR)
		break

	}
}

func (s *Scanner) advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

func (s *Scanner) s.addToken(tokenType TokenType) {

}

func (s *Scanner) s.addTokenLiteral(tokenType TokenType, literal any) {
	input := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{tokenType: tokenType, literal: literal, lexeme: input, line: s.line})
}
