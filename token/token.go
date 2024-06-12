package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	IDENTIFIER = "IDENTIFIER" //variable and function names
	INT        = "INT"        // integer, duh
	LITERAL    = "LITERAL"    // for strings and stuff, idk

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	//delimiters
	COMMA       = ","
	EXCLAMATION = "!"

	LBRACKET = "("
	RBRACKET = ")"
	LBRACE   = "{"
	RBRACE   = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
