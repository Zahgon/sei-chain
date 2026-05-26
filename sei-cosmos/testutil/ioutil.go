package testutil

import (
	"io"
	"os"
	"sync/atomic"
	"testing"

	"github.com/spf13/cobra"
)

// BufferReader is implemented by types that read from a string buffer.
type BufferReader interface {
	io.Reader
	Reset(string)
}

// BufferWriter is implemented by types that write to a buffer.
type BufferWriter interface {
	io.Writer
	Reset()
	Bytes() []byte
	String() string
}

// ApplyMockIO replaces stdin/out/err with buffers that can be used during testing.
// Returns an input BufferReader and an output BufferWriter.
func ApplyMockIO(c *cobra.Command) (BufferReader, BufferWriter) {
	_ = "STUB: not implemented"
	return *new(BufferReader), *new(BufferWriter)
}

// ApplyMockIODiscardOutputs replaces a cobra.Command output and error streams with a dummy io.Writer.
// Replaces and returns the io.Reader associated to the cobra.Command input stream.
func ApplyMockIODiscardOutErr(c *cobra.Command) BufferReader {
	_ = "STUB: not implemented"
	return *new(BufferReader)
}

// Write the given string to a new temporary file.
// Returns an open file for the test to use.
func WriteToNewTempFile(t testing.TB, s string) *os.File { _ = "STUB: not implemented"; return nil }

var tmpFileCounter atomic.Int32

// TempFile returns a writable temporary file for the test to use.
func TempFile(t testing.TB) *os.File { _ = "STUB: not implemented"; return nil }
