package errorz

import "fmt"

const (
	ERROR_RED    = "\033[0;31m"
	NORMAL_WHITE = "\u001b[37m"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(line int, msg string) {
	Report(line, "", msg)
}

func Report(line int, where string, msg string) {
	fmt.Printf("%s [line %d] Error%s:  %s %s \n", ERROR_RED, line, where, msg, NORMAL_WHITE)
}
