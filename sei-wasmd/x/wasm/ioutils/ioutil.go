package ioutils

import (
	"io"
)

// Uncompress returns gzip uncompressed content if input was gzip, or original src otherwise
func Uncompress(src []byte, limit uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G115 -- checked above

// LimitReader returns a Reader that reads from r
// but stops with types.ErrLimit after n bytes.
// The underlying implementation is a *io.LimitedReader.
func LimitReader(r io.Reader, n int64) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type LimitedReader struct {
	r *io.LimitedReader
}

func (l *LimitedReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
