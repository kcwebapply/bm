package util

import (
	"fmt"
	"os"
)

// LoggingError prints out  argument message and cancel commands when this method was called.
// I only prepare this function to fix all error message format to one point.
func LoggingError(message string) {
	fmt.Printf(message)
	fmt.Print("\n")
	os.Exit(0)
}
