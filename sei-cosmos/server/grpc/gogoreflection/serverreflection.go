/*
 *
 * Copyright 2016 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

/*
Package reflection implements server reflection service.

The service implemented is defined in:
https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:

	import "google.golang.org/grpc/reflection"

	s := grpc.NewServer()
	pb.RegisterYourOwnServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	s.Serve(lis)
*/
package gogoreflection // import "google.golang.org/grpc/reflection"

import (
	"reflect"
	"sync"

	// nolint: staticcheck

	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/grpc"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1"
)

type serverReflectionServer struct {
	rpb.UnimplementedServerReflectionServer
	s *grpc.Server

	initSymbols  sync.Once
	serviceNames []string
	symbols      map[string]*dpb.FileDescriptorProto // map of fully-qualified names to files
}

// Register registers the server reflection service on the given gRPC server.
func Register(s *grpc.Server) { _ = "STUB: not implemented"; return }

// protoMessage is used for type assertion on proto messages.
// Generated proto message implements function Descriptor(), but Descriptor()
// is not part of interface proto.Message. This interface is needed to
// call Descriptor().
type protoMessage interface {
	Descriptor() ([]byte, []int)
}

func (s *serverReflectionServer) getSymbols() (svcNames []string, symbolIndex map[string]*dpb.FileDescriptorProto) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverReflectionServer) processFile(fd *dpb.FileDescriptorProto, processed map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func (s *serverReflectionServer) processMessage(fd *dpb.FileDescriptorProto, prefix string, msg *dpb.DescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (s *serverReflectionServer) processEnum(fd *dpb.FileDescriptorProto, prefix string, en *dpb.EnumDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (s *serverReflectionServer) processField(fd *dpb.FileDescriptorProto, prefix string, fld *dpb.FieldDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func fqn(prefix, name string) string { _ = "STUB: not implemented"; return "" }

// fileDescForType gets the file descriptor for the given type.
// The given type should be a proto message.
func (s *serverReflectionServer) fileDescForType(st reflect.Type) (*dpb.FileDescriptorProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decodeFileDesc does decompression and unmarshalling on the given
// file descriptor byte slice.
func decodeFileDesc(enc []byte) (*dpb.FileDescriptorProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decompress does gzip decompression.
func decompress(b []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func typeForName(name string) (reflect.Type, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), nil
}

func fileDescContainingExtension(st reflect.Type, ext int32) (*dpb.FileDescriptorProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *serverReflectionServer) allExtensionNumbersForType(st reflect.Type) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fileDescWithDependencies returns a slice of serialized fileDescriptors in
// wire format ([]byte). The fileDescriptors will include fd and all the
// transitive dependencies of fd with names not in sentFileDescriptors.
func fileDescWithDependencies(fd *dpb.FileDescriptorProto, sentFileDescriptors map[string]bool) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fileDescEncodingByFilename finds the file descriptor for given filename,
// finds all of its previously unsent transitive dependencies, does marshalling
// on them, and returns the marshalled result.
func (s *serverReflectionServer) fileDescEncodingByFilename(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseMetadata finds the file descriptor bytes specified meta.
// For SupportPackageIsVersion4, m is the name of the proto file, we
// call proto.FileDescriptor to get the byte slice.
// For SupportPackageIsVersion3, m is a byte slice itself.
func parseMetadata(meta interface{}) ([]byte, bool) {
	_ = "STUB: not implemented"
	// Check if meta is the file name.
	return nil, false
}

// Check if meta is the byte slice.

// fileDescEncodingContainingSymbol finds the file descriptor containing the
// given symbol, finds all of its previously unsent transitive dependencies,
// does marshalling on them, and returns the marshalled result. The given symbol
// can be a type, a service or a method.
func (s *serverReflectionServer) fileDescEncodingContainingSymbol(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if it's a type name that was not present in the
// transitive dependencies of the registered services.

// fileDescEncodingContainingExtension finds the file descriptor containing
// given extension, finds all of its previously unsent transitive dependencies,
// does marshalling on them, and returns the marshalled result.
func (s *serverReflectionServer) fileDescEncodingContainingExtension(typeName string, extNum int32, sentFileDescriptors map[string]bool) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allExtensionNumbersForTypeName returns all extension numbers for the given type.
func (s *serverReflectionServer) allExtensionNumbersForTypeName(name string) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ServerReflectionInfo is the reflection service handler.
func (s *serverReflectionServer) ServerReflectionInfo(stream rpb.ServerReflection_ServerReflectionInfoServer) error {
	_ = "STUB: not implemented"
	return nil
}
