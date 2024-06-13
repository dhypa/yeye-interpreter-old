package lexer

import (
	"interpreterproject/token"
)

type Lexer struct {
	input    string
	position uint64 // points to index which holds ch
	readpos  uint64 // position + 1
	ch       byte   // holds the character being checked
	inputlen uint64
}

func New(input string) *Lexer {

	lexer := &Lexer{input: input, inputlen: (uint64(len(input)))}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readpos >= lexer.inputlen {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readpos]
	}
	lexer.position = lexer.readpos
	lexer.readpos++
}

func (lexer *Lexer) nextToken() token.Token {
	var tok token.Token

	lexer.INHALEWHITESPACE()

	switch lexer.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(lexer.ch)}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(lexer.ch)}
	case '*':
		tok = token.Token{Type: token.MULTIPLY, Literal: string(lexer.ch)}
	case '/':
		tok = token.Token{Type: token.DIVIDE, Literal: string(lexer.ch)}
	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(lexer.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(lexer.ch)}
	case '!':
		tok = token.Token{Type: token.EXCLAMATION, Literal: string(lexer.ch)}
	case '(':
		tok = token.Token{Type: token.LBRACKET, Literal: string(lexer.ch)}
	case ')':
		tok = token.Token{Type: token.RBRACKET, Literal: string(lexer.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(lexer.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(lexer.ch)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}

	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.getIdentifier()
			tok.Type = token.LookUpKeyword(tok.Literal)
			return tok // getIdentifier() already advances the read position by calling readChar()
		} else if isDigit(lexer.ch) {
			tok.Literal = lexer.getInt()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(lexer.ch)}
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) getInt() string {
	firstPosition := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[firstPosition:lexer.position]
}

// returns a substring of a indentifier made up of only letters
func (lexer *Lexer) getIdentifier() string {
	firstPosition := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[firstPosition:lexer.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Letters for purpose of identifier may have underscores
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_')
}

func (lexer *Lexer) INHALEWHITESPACE() {
	for lexer.ch == ' ' || lexer.ch == '\n' || lexer.ch == '\t' || lexer.ch == '\r' {
		lexer.readChar()
	}
}
