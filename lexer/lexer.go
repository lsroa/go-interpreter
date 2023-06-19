package lexer

import (
	"lsroa/go_interpreter/token"
)

type Lexer struct {
	ch           byte
	input        string
	position     int
	readPosition int
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(literal, token.EQ)
		} else {
			tok = token.New(l.ch, token.ASSIGN)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.New(literal, token.NOT_EQ)
		} else {
			tok = token.New(l.ch, token.BANG)
		}
	case '-':
		tok = token.New(l.ch, token.MINUS)
	case '+':
		tok = token.New(l.ch, token.PLUS)
	case '/':
		tok = token.New(l.ch, token.SLASH)
	case '*':
		tok = token.New(l.ch, token.ASTERISK)
	case '<':
		tok = token.New(l.ch, token.LT)
	case '>':
		tok = token.New(l.ch, token.GT)
	case ',':
		tok = token.New(l.ch, token.COMMA)
	case ';':
		tok = token.New(l.ch, token.SEMICOLON)
	case '{':
		tok = token.New(l.ch, token.L_BRACE)
	case '}':
		tok = token.New(l.ch, token.R_BRACE)
	case '(':
		tok = token.New(l.ch, token.L_PAREN)
	case ')':
		tok = token.New(l.ch, token.R_PAREN)
	case '[':
		tok = token.New(l.ch, token.L_SQUARE)
	case ']':
		tok = token.New(l.ch, token.R_SQUARE)
	case 0:
		tok = token.New(l.ch, token.EOF)

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Kind = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Kind = token.INT
			return tok
		} else {
			tok = token.New(l.ch, token.ILLEGAL)
		}
	}

	l.readChar()
	return tok
}
