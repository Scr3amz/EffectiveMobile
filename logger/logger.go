package logger

import (
	"log"
	"os"
)

type Logger struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	WarningLog *log.Logger
}

func NewLogger(fileName string) Logger {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}

	logger := Logger{
		InfoLog:    log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog:   log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarningLog: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
	return logger
}
