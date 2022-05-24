package errhelp

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestLogIfErrorWithNoError(t *testing.T) {
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	action := "testing"

	LogIfError(action, nil)
	s := strings.TrimSpace(logOutput.String())
	if s != "" {
		t.Errorf("Expected log message to be blank, but the message is [%s]\n", s)
	}
}

func TestLogIfErrorWithError(t *testing.T) {
	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)

	action := "testing"
	err := errors.New("Something bad happened")
	expectedMessage := fmt.Sprintf("Error %s: %v", action, err)

	LogIfError(action, err)
	s := strings.TrimSpace(logOutput.String())
	if !strings.HasSuffix(s, expectedMessage) {
		t.Errorf("Expected log message to be [%s], but the message is [%s]\n", expectedMessage, s)
	}
}
