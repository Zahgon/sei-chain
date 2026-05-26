package main

import (
	"flag"
	"iter"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var flags flag.FlagSet

var moduleFlag = flags.String("module", "", "prefix to strip from the absolute generated file path. Same as in protoc-gen-go")

func (d md) GetBoolOption(ext protoreflect.ExtensionTypeDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

type md struct{ protoreflect.MessageDescriptor }
type mds = map[protoreflect.FullName]md

func (d md) walk(yield func(md) bool) bool { _ = "STUB: not implemented"; return false }

func allMDs(files *protoregistry.Files) iter.Seq[md] { _ = "STUB: not implemented"; return nil }

func OrPanic(err error) { _ = "STUB: not implemented"; return }

func OrPanic1[T any](v T, err error) T {
	_ = "STUB: not implemented"
	return *

	// run reads the proto descriptors and checks that the hashable messages satisfy the following constraints:
	// * all hashable messages have to use proto3 syntax
	// * message fields of hashable messages have to be hashable as well
	// * fields of hashable messages have to be repeated/optional (explicit presence)
	// * fields of hashable messages cannot be maps
	new(T)
}

func run(p *protogen.Plugin) error { _ = "STUB: not implemented"; return nil }

// Re-unmarshal proto files, so that dynamic options are registered.

type pm struct{ *protogen.Message }

func (m pm) walk(yield func(pm) bool) bool { _ = "STUB: not implemented"; return false }

func allPMs(f *protogen.File) iter.Seq[pm] { _ = "STUB: not implemented"; return nil }

func generateHashableFiles(p *protogen.Plugin, descs mds) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	protogen.Options{ParamFunc: flags.Set}.Run(run)
}
