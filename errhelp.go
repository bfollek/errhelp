// Package errhelp provides simple helper functions to make golang error handling easier.
package errhelp

import (
	"fmt"
	"log"
)

// LogIfError - If `err` is not null, builds a message and logs it.
// `action` is a string describing what was happening, e.g. "opening history file".
func LogIfError(action string, err error) {
	if err != nil {
		log.Println(buildMsg(action, err))
	}
}

// FatalIfError - If `err` is not null, builds a message, logs it, and stops the program.
// `action` is a string describing what was happening, e.g. "opening history file".
func FatalIfError(action string, err error) {
	if err != nil {
		log.Fatal(buildMsg(action, err))
	}
}

func buildMsg(action string, err error) string {
	return fmt.Sprintf("Error %s: %v", action, err)
}
