package main

import (
	"interpreter/lox"
)

func main() {
	l := lox.NewLox()
	l.RunPrompt()
}
