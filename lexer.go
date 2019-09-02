package graphql_lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	value        byte
}

func NewLexer(input string) Lexer {
	l := Lexer{input: input}
	l.advance()
	return l
}

func (l *Lexer) advance() {
	if l.readPosition >= len(l.input) {
		l.value = 0
	} else {
		l.value = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()

	switch l.value {
	case ',':
		tok = readToken(COMMA, l.value, l.position, l.readPosition)
	case ':':
		tok = readToken(COLON, l.value, l.position, l.readPosition)
	case '(':
		tok = readToken(LPAREN, l.value, l.position, l.readPosition)
	case ')':
		tok = readToken(RPAREN, l.value, l.position, l.readPosition)
	case '{':
		tok = readToken(LBRACE, l.value, l.position, l.readPosition)
	case '}':
		tok = readToken(RBRACE, l.value, l.position, l.readPosition)
	case '[':
		tok = readToken(LBRACKET, l.value, l.position, l.readPosition)
	case ']':
		tok = readToken(RBRACKET, l.value, l.position, l.readPosition)
	case '=':
		tok = readToken(EQUALS, l.value, l.position, l.readPosition)
	case '!':
		tok = readToken(BANG, l.value, l.position, l.readPosition)
	case '$':
		tok = readToken(DOLLAR, l.value, l.position, l.readPosition)
	case '@':
		tok = readToken(AT, l.value, l.position, l.readPosition)
	case '|':
		tok = readToken(PIPE, l.value, l.position, l.readPosition)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isStringLiteral(l.value) {
			tok.Position = l.position
			tok.Literal = l.readStringLiteral()
			tok.Type = STRING
			tok.ReadPosition = l.readPosition
			return tok
		} else if isAlphabet(l.value) {
			tok.Position = l.position
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdentifier(tok.Literal)
			tok.ReadPosition = l.readPosition
			return tok
		} else if isDigit(l.value) {
			tok.Position = l.position
			tok.Type = INT
			tok.Literal = l.readNumber()
			tok.ReadPosition = l.readPosition
			return tok
		} else {
			// TODO: Implement BLOCKSTRING, DIRECTIVE cases
			tok = readToken(ILLEGAL, l.value, l.position, l.readPosition)
		}
	}

	l.advance()
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isAlphabet(l.value) {
		l.advance()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.value) {
		l.advance()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readStringLiteral() string {
	l.advance()
	pos := l.position
	for !isStringLiteral(l.value) {
		l.advance()
	}
	l.advance()
	return l.input[pos : l.position-1]
}

func (l *Lexer) skipWhitespace() {
	for l.value == ' ' || l.value == '\t' || l.value == '\n' || l.value == '\r' {
		l.advance()
	}
}

func readToken(t TokenType, l byte, p int, rp int) Token {
	return Token{
		Literal:      string(l),
		Type:         t,
		Position:     p,
		ReadPosition: rp,
	}
}

func isAlphabet(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isStringLiteral(c byte) bool {
	return c == '"'
}
