package errors

import (
	"log"
	"os"
)

// ErrorMsg print the error message to the screen
func ErrorMsg(err string) {
	log.Println(err)
	os.Exit(1)
}
