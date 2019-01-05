package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	// Info :
	Info *log.Logger
	// Error :
	Error *log.Logger
)

func init() {
	file, err := os.OpenFile("data/logs/combined.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Log error init error", err)
	}

	infoLogger := io.MultiWriter(file, os.Stdout)
	errorLogger := io.MultiWriter(file, os.Stderr)

	set(infoLogger, errorLogger)
}

func set(infoHandle io.Writer, errorHandle io.Writer) {
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
