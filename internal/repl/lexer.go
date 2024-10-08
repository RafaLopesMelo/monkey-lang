package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RafaLopesMelo/monkey-lang/internal/lexer"
	"github.com/RafaLopesMelo/monkey-lang/internal/token"
)

func StartLexerRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}
}
