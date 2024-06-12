package token

type Token struct {
	Type    string
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	IDENTIFIER = "IDENTIFIER" //variable and function names
	INT        = "INT"        // integer, duh
	LITERAL    = "LITERAL"    // for strings and stuff, idk

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

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
