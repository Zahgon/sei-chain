package snapshots

import (
	"bufio"
	"compress/zlib"
	"io"

	protoio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
)

const (
	// Do not change chunk size without new snapshot format (must be uniform across nodes)
	snapshotChunkSize  = uint64(10e6)
	snapshotBufferSize = int(snapshotChunkSize)
	// Do not change compression level without new snapshot format (must be uniform across nodes)
	snapshotCompressionLevel = 7
)

// StreamWriter set up a stream pipeline to serialize snapshot nodes:
// Exported Items -> delimited Protobuf -> zlib -> buffer -> chunkWriter -> chan io.ReadCloser
type StreamWriter struct {
	chunkWriter *ChunkWriter
	bufWriter   *bufio.Writer
	zWriter     *zlib.Writer
	protoWriter protoio.WriteCloser
}

// NewStreamWriter set up a stream pipeline to serialize snapshot DB records.
func NewStreamWriter(ch chan<- io.ReadCloser) *StreamWriter { _ = "STUB: not implemented"; return nil }

// WriteMsg implements protoio.Write interface
func (sw *StreamWriter) WriteMsg(msg proto.Message) error { _ = "STUB: not implemented"; return nil }

// Close implements io.Closer interface
func (sw *StreamWriter) Close() error { _ = "STUB: not implemented"; return nil }

// CloseWithError pass error to chunkWriter
func (sw *StreamWriter) CloseWithError(err error) { _ = "STUB: not implemented"; return }

// StreamReader set up a restore stream pipeline
// chan io.ReadCloser -> chunkReader -> zlib -> delimited Protobuf -> ExportNode
type StreamReader struct {
	chunkReader *ChunkReader
	zReader     io.ReadCloser
	protoReader protoio.ReadCloser
}

// NewStreamReader set up a restore stream pipeline.
func NewStreamReader(chunks <-chan io.ReadCloser) (*StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadMsg implements protoio.Reader interface
func (sr *StreamReader) ReadMsg(msg proto.Message) error { _ = "STUB: not implemented"; return nil }

// Close implements io.Closer interface
func (sr *StreamReader) Close() error { _ = "STUB: not implemented"; return nil }
