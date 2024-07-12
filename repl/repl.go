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

		l := lexer.New(scanner.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
