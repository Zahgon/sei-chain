package input

import (
	"bufio"
	"io"
)

// MinPassLength is the minimum acceptable password length
const MinPassLength = 8

// GetPassword will prompt for a password one-time (to sign a tx)
// It enforces the password length
func GetPassword(prompt string, buf *bufio.Reader) (pass string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Return the given password to the upstream client so it can handle a
// non-STDIN failure gracefully.

// GetConfirmation will request user give the confirmation from stdin.
// "y", "Y", "yes", "YES", and "Yes" all count as confirmations.
// If the input is not recognized, it returns false and a nil error.
func GetConfirmation(prompt string, r *bufio.Reader, w io.Writer) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetString simply returns the trimmed string output of a given reader.
func GetString(prompt string, buf *bufio.Reader) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// inputIsTty returns true iff we have an interactive prompt,
// where we can disable echo and request to repeat the password.
// If false, we can optimize for piped input from another command
func inputIsTty() bool { _ = "STUB: not implemented"; return false }

// readLineFromBuf reads one line from reader.
// Subsequent calls reuse the same buffer, so we don't lose
// any input when reading a password twice (to verify)
func readLineFromBuf(buf *bufio.Reader) (string, error) { _ = "STUB: not implemented"; return "", nil }

// If by any chance the error is EOF, but we were actually able to read
// something from the reader then don't return the EOF error.
// If we didn't read anything from the reader and got the EOF error, then
// it's safe to return EOF back to the caller.

// exit the switch statement
