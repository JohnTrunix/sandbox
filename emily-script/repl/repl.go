package repl

import (
	"bufio"
	"emily-script/lexer"
	"emily-script/token"
	"fmt"
	"io"
	"os"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		} else if scanner.Text() == "exit" {
			fmt.Printf("See you!")
			os.Exit(0)
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
