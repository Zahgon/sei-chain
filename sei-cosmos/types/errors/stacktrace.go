package errors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func matchesFunc(f errors.Frame, prefixes ...string) bool { _ = "STUB: not implemented"; return false }

// funcName returns the name of this function, if known.
func funcName(f errors.Frame) string {
	_ = "STUB: not implemented"
	// this looks a bit like magic, but follows example here:
	// https://github.com/pkg/errors/blob/v0.8.1/stack.go#L43-L50
	return ""
}

func fileLine(f errors.Frame) (string, int) {
	_ = "STUB: not implemented"
	// this looks a bit like magic, but follows example here:
	// https://github.com/pkg/errors/blob/v0.8.1/stack.go#L14-L27
	// as this is where we get the Frames
	return "", 0
}

func trimInternal(st errors.StackTrace) errors.StackTrace {
	_ = "STUB: not implemented"
	// trim our internal parts here
	// manual error creation, or runtime for caught panics
	return *new(errors.StackTrace)
}

// where we create errors

// runtime are added on panics

// _test is defined in coverage tests, causing failure
// "/_test/"

// trim out outer wrappers (runtime.goexit and test library if present)

func writeSimpleFrame(s io.Writer, f errors.Frame) { _ = "STUB: not implemented"; return }

// cut file at "github.com/"
// TODO: generalize better for other hosts?

// Format works like pkg/errors, with additions.
// %s is just the error message
// %+v is the full stack trace
// %v appends a compressed [filename:line] where the error
//
//	was created
//
// Inspired by https://github.com/pkg/errors/blob/v0.8.1/errors.go#L162-L176
func (e *wrappedError) Format(s fmt.State, verb rune) {
	_ = "STUB: not implemented"
	// normal output here....
	return
}

// work with the stack trace... whole or part

// stackTrace returns the first found stack trace frame carried by given error
// or any wrapped error. It returns nil if no stack trace is found.
func stackTrace(err error) errors.StackTrace {
	_ = "STUB: not implemented"
	return *new(errors.StackTrace)
}
