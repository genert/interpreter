package token

// Type is defined as string.
// This allows us to use many different values as Types, which in turn allows us to distinguish
// between different types of tokens.
// Using string also has the advantage of being easy to debug
// without a lot of boilerplate and helper functions: we can just print a string.
type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL signifies a token/character we don’t know about.
	ILLEGAL = "ILLEGAL"

	// EOF stands for “end of file”, which tells our parser later on that it can stop.
	EOF = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"function": FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
