package serror

import (
	"errors"
	"fmt"
	"runtime"
)

type sourceError struct {
	error  error
	source string
}

func (e *sourceError) Error() string {
	return e.error.Error() + "\n\tat " + e.source
}

func (e *sourceError) Unwrap() error {
	return e.error
}

// Wrap wraps an existing error by source information.
func Wrap(err error) error {
	return &sourceError{
		error:  err,
		source: source(3),
	}
}

// Error creates an error with source information.
func Error(text string) error {
	return &sourceError{
		error:  errors.New(text),
		source: source(3),
	}
}

// Errorf creates an error with source information.
// It takes the same arguments as fmt.Errorf.
func Errorf(format string, args ...any) error {
	return &sourceError{
		error:  fmt.Errorf(format, args...),
		source: source(3),
	}
}

func source(skip int) string {
	var pcs [1]uintptr
	runtime.Callers(skip, pcs[:])
	fs := runtime.CallersFrames([]uintptr{pcs[0]})
	f, _ := fs.Next()
	return fmt.Sprintf("%s:%d", f.File, f.Line)
}
