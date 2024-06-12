package lexer

import (
	"interpreterproject/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let topSpeed = 14!
	let distanceToHome = 1367!

	let timeTaken = fn(distance, speed) {
		distance / speed!
	}!

	let result = timeTaken(distance, speed)!
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, // first line
		{token.IDENTIFIER, "topSpeed"},
		{token.ASSIGN, "="},
		{token.INT, "14"},
		{token.EXCLAMATION, "!"},
		{token.LET, "let"}, // second line
		{token.IDENTIFIER, "distanceToHome"},
		{token.ASSIGN, "="},
		{token.INT, "1367"},
		{token.EXCLAMATION, "!"},
		{token.LET, "let"}, // fourth line
		{token.IDENTIFIER, "timeTaken"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LBRACKET, "("},
		{token.IDENTIFIER, "distance"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "speed"},
		{token.RBRACKET, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "distance"}, // line 5
		{token.DIVIDE, "/"},
		{token.IDENTIFIER, "speed"},
		{token.EXCLAMATION, "!"},
		{token.RBRACE, "}"}, //line 6
		{token.EXCLAMATION, "!"},
		{token.LET, "let"}, //line 7
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "timeTaken"},
		{token.LBRACKET, "("},
		{token.IDENTIFIER, "distance"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "speed"},
		{token.RBRACKET, ")"},
		{token.EXCLAMATION, "!"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.nextToken()

		if tok.Type != tt.expectedType {

		}

	}

}
