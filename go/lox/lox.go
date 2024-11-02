package lox

import (
	"bufio"
	"fmt"
	astprinter "interpreter/astPrinter"
	"interpreter/errorz"
	"interpreter/parser"
	"os"
)

type Lox struct {
	hadError bool
}

func NewLox() Lox {
	return Lox{hadError: false}
}

func (l *Lox) RunFile(path string) {
	data, err := os.ReadFile(path)
	errorz.CheckError(err)
	run(string(data))
	if l.hadError {
		os.Exit(1)
	}
}

func (l *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(">")
		text, err := reader.ReadString('\n')
		errorz.CheckError(err)
		run(text)
		l.hadError = false
	}

}

func run(source string) {
	sc := NewScanner(source)
	tokens := sc.scanTokens()
	// for ind, t := range tokens {
	// 	fmt.Printf("Token[%d]: %s \n", ind, t.ToString())
	// }
	parser := parser.NewParser(tokens)
	expression, err := parser.Parse()
	if err != nil {
		return
	}
	printer := astprinter.NewPrinter()
	fmt.Println("Printing expressions...")
	printer.Print(&expression)
}
