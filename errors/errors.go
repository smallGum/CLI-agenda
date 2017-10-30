package errors

import (
	"fmt"
	"os"
)

// ErrorMsg print the error message to the screen
func ErrorMsg(err string) {
	fmt.Println(err)
	os.Exit(1)
}
