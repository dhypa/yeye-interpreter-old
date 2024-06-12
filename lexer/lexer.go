package lexer

import "interpreterproject/token"

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

	switch lexer.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(lexer.ch)}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(lexer.ch)}
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
		tok = token.Token{Type: token.LBRACE, Literal: string(lexer.ch)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	}

	lexer.readChar()
	return tok
}
