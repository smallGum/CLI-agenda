package errors

import (
	"log"
	"os"
)

// ErrorMsg record the error messsage and print the error message
func ErrorMsg(usr, err string) {
	file, _ := os.OpenFile("log_files/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0)
	logger := log.New(file, usr+": ", log.Lshortfile)
	logger.Println(err)
	log.Println(err)

	file.Close()
}

// LogMeetingOperation record the operations about meetings and print the feedback
func LogMeetingOperation(usr, oper string) {
	file, _ := os.OpenFile("log_files/meetings.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0)
	logger := log.New(file, usr+": ", log.Lshortfile)
	logger.Println(oper)
	log.Println(oper)

	file.Close()
}
