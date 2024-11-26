// gohello/internal/utils/utils.go

package utils

import (
	"fmt"
	"os"
)

func Print(str string) {
	fmt.Fprintln(os.Stdout, str)
}

func PrintMsg(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func PrintErr(errMsg string) {
	fmt.Fprintln(os.Stderr, errMsg)
}
