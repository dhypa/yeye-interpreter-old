package lexer

type Lexer struct {
	input    string
	position uint64 // points to index which holds ch
	readpos  uint64 // position + 1
	ch       byte   // holds the character being checked
	inputlen uint64
}

func New(input string) *Lexer {

	lexer := &Lexer{input: input, inputlen: (uint64(len(input)))}

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
