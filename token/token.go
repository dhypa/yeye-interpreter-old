package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookUpKeyword(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}
	return IDENTIFIER

}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	IDENTIFIER = "IDENTIFIER" //variable and function names
	INT        = "INT"        // integer, duh
	LITERAL    = "LITERAL"    // for strings and stuff, idk

	// operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	MULTIPLY  = "*"
	DIVIDE    = "/"
	MORE_THAN = ">"
	LESS_THAN = "<"
	NOT       = ";"
	EQV       = "=="
	NOT_EQV   = "!="

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
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)
