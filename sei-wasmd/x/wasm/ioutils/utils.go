package ioutils

// Note: []byte can never be const as they are inherently mutable
var (
	// magic bytes to identify gzip.
	// See https://www.ietf.org/rfc/rfc1952.txt
	// and https://github.com/golang/go/blob/master/src/net/http/sniff.go#L186
	gzipIdent = []byte("\x1F\x8B\x08")

	wasmIdent = []byte("\x00\x61\x73\x6D")
)

// IsGzip returns checks if the file contents are gzip compressed
func IsGzip(input []byte) bool { _ = "STUB: not implemented"; return false }

// IsWasm checks if the file contents are of wasm binary
func IsWasm(input []byte) bool { _ = "STUB: not implemented"; return false }

// GzipIt compresses the input ([]byte)
func GzipIt(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Create gzip writer.
	return nil, nil
}

// You must close this first to flush the bytes to the buffer.
