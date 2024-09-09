package utils

import (
	"log"
	"os"
)

// Logger is a centralized logging function
func Logger() *log.Logger {
	return log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
