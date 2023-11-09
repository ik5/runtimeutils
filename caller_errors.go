package runtimeutils

import "errors"

var (
	ErrNoFramesWereFound    = errors.New("no caller frames were provided")
	ErrNoCallerInfoWasFound = errors.New("no caller information was found")
)
