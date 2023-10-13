package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init(logFilePath string) {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
	Logger = log.New(logFile, "Backend: ", log.Ldate|log.Ltime)
}
