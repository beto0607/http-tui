package logger

import (
	"log"
	"os"
)

var logFile *os.File

func NewLogger() *log.Logger {
	logFile, err := os.OpenFile("tmp/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	newLogger := log.New(logFile, "http-tui", log.Lshortfile)

	return newLogger
}

func StopLogger() {
	logFile.Close()
}
