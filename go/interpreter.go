package main

import (
	"fmt"
	"interpreter/lox"
	"os"
)

func main() {
	for _, val := range os.Args {
		fmt.Println(val)
	}
	l := lox.NewLox()
	if len(os.Args) > 1 {
		l.RunFile(os.Args[1])
	} else {
		l.RunPrompt()
	}
}
