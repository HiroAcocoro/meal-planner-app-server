package errors

import (
	"log"
)

// Log Fatal Errors allows for error handling that terminates the app
// making this reusable allows for future logging operations
func LogFatalError(err error) {
	log.Fatal(err)
}

func LogError(err error) {
	log.Println(err)
}
