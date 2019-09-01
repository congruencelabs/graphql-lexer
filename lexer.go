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
		tok = readToken(COMMA, l.value)
	case ':':
		tok = readToken(COLON, l.value)
	case '(':
		tok = readToken(LPAREN, l.value)
	case ')':
		tok = readToken(RPAREN, l.value)
	case '{':
		tok = readToken(LBRACE, l.value)
	case '}':
		tok = readToken(RBRACE, l.value)
	case '[':
		tok = readToken(LBRACKET, l.value)
	case ']':
		tok = readToken(RBRACKET, l.value)
	case '=':
		tok = readToken(EQUALS, l.value)
	case '!':
		tok = readToken(BANG, l.value)
	case '$':
		tok = readToken(DOLLAR, l.value)
	case '@':
		tok = readToken(AT, l.value)
	case '|':
		tok = readToken(PIPE, l.value)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isStringLiteral(l.value) {
			tok.Literal = l.readStringLiteral()
			tok.Type = STRING
			return tok
		} else if isAlphabet(l.value) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.value) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// TODO: Implement BLOCKSTRING, DIRECTIVE cases
			tok = readToken(ILLEGAL, l.value)
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
	return l.input[pos:l.position - 1]
}

func (l *Lexer) skipWhitespace() {
	for l.value == ' ' || l.value == '\t' || l.value == '\n' || l.value == '\r' {
		l.advance()
	}
}

func readToken(tokenType TokenType, literal byte) Token {
	return Token{Literal: string(literal), Type: tokenType}
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
