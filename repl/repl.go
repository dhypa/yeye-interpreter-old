package repl

import (
	"bufio"
	"fmt"
	"interpreterproject/lexer"
	"interpreterproject/token"
	"os"

	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")

		if !scanner.Scan() {
			return
		}

		lexer := lexer.New(scanner.Text())
		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
