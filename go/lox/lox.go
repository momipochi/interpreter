package lox

import (
	"bufio"
	"fmt"
	astprinter "interpreter/astPrinter"
	"interpreter/errorz"
	"interpreter/interpreter"
	"interpreter/parser"
	"os"
)

type Lox struct {
	hadError        bool
	hadRuntimeError bool
	interpreter     *interpreter.Interpreter
}

func NewLox() *Lox {
	tmp := &Lox{hadError: false, hadRuntimeError: false}
	tmp.interpreter = interpreter.NewInterpreter(&tmp.hadError, &tmp.hadRuntimeError)
	return tmp
}

func (l *Lox) RunFile(path string) {
	data, err := os.ReadFile(path)
	errorz.CheckError(err)
	l.run(string(data))
	if l.hadError {
		os.Exit(1)
	}
	if l.hadRuntimeError {
		os.Exit(1)
	}
}

func (l *Lox) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(">")
		text, err := reader.ReadString('\n')
		errorz.CheckError(err)
		l.run(text)
		l.hadError = false
	}

}

func (l *Lox) run(source string) {
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
	fmt.Println("\nInterpreted value:")
	l.interpreter.Interpret(&expression)
	fmt.Println("\nPrinting expressions...")

	printer.Print(&expression)
}
