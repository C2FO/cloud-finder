package logging

import (
	"log"
	"os"
)

// debug is whether or not debug logging is enabled. By Default is is false,
// since the output of the program will be mainly eval'd by other processes.
var debug bool

// EnableDebug enables debug logging.
func EnableDebug() {
	debug = true
}

// Printf wraps the standard lib's log.Printf to only print if debug is enabled.
func Printf(format string, v ...interface{}) {
	if debug {
		log.Printf(format, v...)
	}
}

// Fatalf wraps the standard lib's log.Fatalf to only print if debug is enabled.
// Otherwise it just call's os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	if debug {
		log.Fatalf(format, v...)
	}
	os.Exit(1)
}
