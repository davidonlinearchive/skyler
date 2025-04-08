package repl

import (
	"bufio"
	"fmt"
	"io"
	"skyler/lexer"
)



const PROMPT = ">>"


func Start( in io.Reader){
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for tok
	}
}