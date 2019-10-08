package errhelp

import "log"

// FatalIfError builds a message and calls log.Fatalf if err is not nil.
// action is a string describing what was happening, e.g. "opening history file".
func FatalIfError(action string, err error) {
	if err != nil {
		log.Fatalf("Error %s: %v", action, err)
	}
}

// LogIfError builds a message and calls log.Printf if err is not null.
// action is a string describing what was happening, e.g. "opening history file".
// Returns true if err is not nil, else false.
func LogIfError(action string, err error) bool {
	if err != nil {
		log.Printf("Error %s: %v", action, err)
		return true
	}
	return false
}
