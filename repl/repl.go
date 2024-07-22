package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/behouba/monkeylang/token"

	"github.com/behouba/monkeylang/lexer"
)

const PROMPT = "[monkey REPL]$ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			return
		}

		l := lexer.New(line)

		for tok := l.NexToken(); tok.Type != token.EOF; tok = l.NexToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
