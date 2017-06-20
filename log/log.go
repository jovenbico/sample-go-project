package log

import (
	"io"
	"log"
	"os"

	"gitlab.com/parkby/parkby-service/config"
)

var (
	// Debug debug logger
	Debug *log.Logger
	// Info info logger
	Info *log.Logger
	// Warn warn logger
	Warn *log.Logger
	// Error error logger
	Error *errorLogger
)

type errorLogger struct {
	*log.Logger
}

func init() {
	filePath := config.LoggerPath
	println(">>> ", filePath)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(multiWriter, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(multiWriter, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = &errorLogger{log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)}
}

func (errLogger *errorLogger) PrintlnErr(err error) {
	if err != nil {
		errLogger.Println(err)
	}
}
