// Package errhelp provides simple helper functions to make golang error handling easier.
package errhelp

import (
	"fmt"
	"io"
	"log"
	"os"
)

// LogIfError - If `err` is not null, build a message and log it to the
// standard golang log. `info` is for any information you want in the log,
// e.g. "creating user".
func LogIfError(info string, err error) {
	LogIfErrorNS(log.Writer(), info, err)
}

// LogIfErrorNS - Like `LogIfError` but for non-standard logging. Use
// this to log to something other than the standard golang log.
// `w` is an io.Writer for your log.
func LogIfErrorNS(w io.Writer, info string, err error) {
	if err != nil {
		fmt.Fprintln(w, buildMsg(info, err))
	}
}

// FatalIfError - If `err` is not null, build a message and log it to the
// standard golang log, then stop the program. `info` is for any information
// you want in the log, e.g. "creating user".
func FatalIfError(info string, err error) {
	FatalIfErrorNS(log.Writer(), info, err)
}

// FatalIfErrorNS - Like `FatalIfError` but for non-standard logging. Use
// this to log to something other than the standard golang log.
// `w` is an io.Writer for your log.
func FatalIfErrorNS(w io.Writer, info string, err error) {
	if err != nil {
		fmt.Fprintln(w, buildMsg(info, err))
		os.Exit(1) // What log.Fatal() does.
	}
}

func buildMsg(info string, err error) string {
	return fmt.Sprintf("Error %s: %v", info, err)
}
