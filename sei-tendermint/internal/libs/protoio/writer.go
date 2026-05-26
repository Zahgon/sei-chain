// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Modified from original GoGo Protobuf to return number of bytes written.

package protoio

import (
	"bytes"
	"io"
	"sync"

	"github.com/gogo/protobuf/proto"
)

// NewDelimitedWriter writes a varint-delimited Protobuf message to a writer. It is
// equivalent to the gogoproto NewDelimitedWriter, except WriteMsg() also returns the
// number of bytes written, which is necessary in the p2p package.
func NewDelimitedWriter(w io.Writer) WriteCloser {
	_ = "STUB: not implemented"
	return *new(WriteCloser)
}

type varintWriter struct {
	w      io.Writer
	lenBuf []byte
	buffer []byte
}

func (w *varintWriter) WriteMsg(msg proto.Message) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:gosec // n is a proto message size, always non-negative

// fallback

//nolint:gosec // len() is always non-negative

func (w *varintWriter) Close() error { _ = "STUB: not implemented"; return nil }

func varintWrittenBytes(m marshaler, size int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // size is a proto message size, always non-negative

var bufPool = &sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func MarshalDelimited(msg proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	// The goal here is to write proto message as is knowning already if
	// the exact size can be retrieved and if so just use that.
	return nil, nil
}

// Otherwise, go down the route of using proto.Marshal,
// and use the buffer pool to retrieve a writer.

// Given that we are reusing buffers, we should
// make a copy of the returned bytes.
