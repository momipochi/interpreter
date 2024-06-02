package lox

import (
	"bufio"
	"fmt"
	"interpreter/errorz"
	"os"
)

type Lox struct {
	hadError bool
}

func NewLox() Lox {
	return Lox{hadError: false}
}

func (l *Lox) runFile(path string) {
	data, err := os.ReadFile(path)
	errorz.CheckError(err)
	content := make([]byte, len(data))
	run(string(content))

	if l.hadError {
		os.Exit(1)
	}
}

func (l *Lox) runPrompt() {
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
	fmt.Println("Func Run to be implemented")
}
