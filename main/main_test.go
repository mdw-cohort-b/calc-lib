package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/mdw-smarty-cohort-b-2024-10/calc-lib"
)

func assertErr(t *testing.T, actual error, targets ...error) {
	for _, target := range targets {
		if !errors.Is(actual, target) {
			t.Helper()
			t.Errorf("want: %v, got: %v", target, actual)
		}
	}
}

func TestHandler_TwoArgsRequired(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	assertErr(t, err, errWrongArgCount)
}
func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "42"})
	assertErr(t, err, errInvalidArgument)
}
func TestHandler_SecondArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"42", "INVALID"})
	assertErr(t, err, errInvalidArgument)
}
func TestHandler_ResultWrittenToOutput(t *testing.T) {
	stdout := &bytes.Buffer{}
	handler := NewHandler(stdout, &calc.Addition{})

	err := handler.Handle([]string{"2", "3"})

	assertErr(t, err, nil)
	if stdout.String() != "5" {
		t.Errorf("want: %v, got: %v", "5", stdout.String())
	}
}
func TestHandler_OutputWriterError(t *testing.T) {
	boink := errors.New("boink")
	stdout := &ErringWriter{err: boink}
	handler := NewHandler(stdout, &calc.Addition{})

	err := handler.Handle([]string{"2", "3"})

	assertErr(t, err, errOutputFailure, boink)
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write([]byte) (n int, err error) {
	return 0, this.err
}
