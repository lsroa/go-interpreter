package token

type Kind string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	GT       = ">"
	LT       = "<"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
	EQ       = "=="
	NOT_EQ   = "!="
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	L_PAREN   = "("
	R_PAREN   = ")"
	L_BRACE   = "{"
	R_BRACE   = "}"
	L_SQUARE  = "["
	R_SQUARE  = "]"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
)

var keywords = map[string]Kind{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) Kind {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Literal string
	Kind
}

func New[T string | byte](literal T, kind Kind) Token {
	return Token{Kind: kind, Literal: string(literal)}
}
