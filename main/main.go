package main

import (
	"fmt"
	"interpreterproject/repl"
	"os"
)

func main() {
	fmt.Println("Enter commands here:")
	repl.Start(os.Stdin, os.Stdout)

}
