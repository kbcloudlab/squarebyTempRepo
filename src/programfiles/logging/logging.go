package logging

import "log"

type Logging struct {
}

func NewLogging() *Logging {
	return &Logging{}
}

func (_l *Logging) ErrorLogPanic(_msg string, _err error) {
	log.Println("[LOG] ", _msg, _err)
	panic("")
}

func (_l *Logging) ErrorLog(_msg string, _err error) {
	log.Println("[LOG] ", _msg, _err)
}

func (_l *Logging) MessageLog(_msg string) {
	log.Println("[LOG] ", _msg)
}
