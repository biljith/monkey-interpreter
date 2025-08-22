package lexer

import "interpreter/token"

type Lexer struct {
	input            string
	position         int
	readPosition     int
	currentCharacter byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentCharacter = 0
	} else {
		l.currentCharacter = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentCharacter) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentCharacter {
	case '=':
		tok = newToken(token.ASSIGN, l.currentCharacter)
	case '+':
		tok = newToken(token.PLUS, l.currentCharacter)
	case '(':
		tok = newToken(token.LPAREN, l.currentCharacter)
	case ')':
		tok = newToken(token.RPAREN, l.currentCharacter)
	case '{':
		tok = newToken(token.LBRACE, l.currentCharacter)
	case '}':
		tok = newToken(token.RBRACE, l.currentCharacter)
	case ',':
		tok = newToken(token.COMMA, l.currentCharacter)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentCharacter)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.currentCharacter) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.currentCharacter) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentCharacter)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhiteSpace(l.currentCharacter) {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentCharacter) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isWhiteSpace(character byte) bool {
	return character == ' ' || character == '\t' || character == '\n' || character == '\r'
}

func isDigit(character byte) bool {
	return character >= '0' && character <= '9'
}

func isLetter(character byte) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || character == '_'
}
