package snapshots

import (
	"io"
)

// ChunkWriter reads an input stream, splits it into fixed-size chunks, and writes them to a
// sequence of io.ReadClosers via a channel.
type ChunkWriter struct {
	ch        chan<- io.ReadCloser
	pipe      *io.PipeWriter
	chunkSize uint64
	written   uint64
	closed    bool
}

// NewChunkWriter creates a new ChunkWriter. If chunkSize is 0, no chunking will be done.
func NewChunkWriter(ch chan<- io.ReadCloser, chunkSize uint64) *ChunkWriter {
	_ = "STUB: not implemented"
	return nil
}

// chunk creates a new chunk.
func (w *ChunkWriter) chunk() error { _ = "STUB: not implemented"; return nil }

// Close implements io.Closer.
func (w *ChunkWriter) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithError closes the writer and sends an error to the reader.
func (w *ChunkWriter) CloseWithError(err error) { _ = "STUB: not implemented"; return }

// create a dummy pipe just to propagate the error to the reader, it always returns nil

// Write implements io.Writer.
func (w *ChunkWriter) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:gosec

//nolint:gosec

// ChunkReader reads chunks from a channel of io.ReadClosers and outputs them as an io.Reader
type ChunkReader struct {
	ch     <-chan io.ReadCloser
	reader io.ReadCloser
}

// NewChunkReader creates a new ChunkReader.
func NewChunkReader(ch <-chan io.ReadCloser) *ChunkReader { _ = "STUB: not implemented"; return nil }

// next fetches the next chunk from the channel, or returns io.EOF if there are no more chunks.
func (r *ChunkReader) next() error { _ = "STUB: not implemented"; return nil }

// Close implements io.ReadCloser.
func (r *ChunkReader) Close() error { _ = "STUB: not implemented"; return nil }

// Read implements io.Reader.
func (r *ChunkReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// DrainChunks drains and closes all remaining chunks from a chunk channel.
func DrainChunks(chunks <-chan io.ReadCloser) { _ = "STUB: not implemented"; return }
