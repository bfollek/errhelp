package errhelp

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestLogIfErrorWithError(t *testing.T) {
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	err := errors.New("File not found")

	expectedResult := true
	result := LogIfError("opening file", err)
	if expectedResult != result {
		t.Errorf("Expected [%t], got [%t]\n", expectedResult, result)
	}
	s := strings.TrimSpace(logOutput.String())
	expectedMessage := "Error opening file: File not found"
	if !strings.HasSuffix(s, expectedMessage) {
		t.Errorf("Expected log message to end with [%s], but the message is [%s]\n", expectedMessage, s)
	}
}

func TestLogIfErrorWithNoError(t *testing.T) {
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	expectedResult := false
	result := LogIfError("opening file", nil)
	if expectedResult != result {
		t.Errorf("Expected [%t], got [%t]\n", expectedResult, result)
	}
	s := strings.TrimSpace(logOutput.String())
	if s != "" {
		t.Errorf("Expected log message to be blank, but the message is [%s]\n", s)
	}
}
