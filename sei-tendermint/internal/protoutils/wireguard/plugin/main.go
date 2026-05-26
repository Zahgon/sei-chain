// Package main implements a protoc plugin that turns wireguard.proto field
// annotations into *wireguard.Schema variables, one per annotated message
// type. The output sits next to the .pb.go files as `<name>.wireguard.go`.
//
// Only one annotation is needed:
//
//	(wireguard.max_count) = N      // cap on a repeated field's occurrences
//
// Descent into nested message fields is automatic: if a field's target type
// has a Schema (i.e. has annotations somewhere in its reachable subtree),
// the parent's rule descends into it. Fields whose target type has no
// annotations are walked past.
//
// Strict mode (`--strict`): every reachable repeated field must carry
// (wireguard.max_count); a missing annotation is a codegen error. Default
// off so this plugin can land before the full audit of repeated fields
// across the proto tree.
//
// TODO: dedup with sei-tendermint/internal/hashable/plugin in a later PR.
package main

import (
	"flag"
	"iter"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var flags flag.FlagSet

var (
	moduleFlag = flags.String("module", "", "prefix to strip from the absolute generated file path. Same as in protoc-gen-go")
	strictFlag = flags.Bool("strict", false, "every reachable repeated field must carry (wireguard.max_count); a missing annotation is a codegen error")
)

func main() {
	protogen.Options{ParamFunc: flags.Set}.Run(run)
}

const (
	wireguardRuntime = "github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils/wireguard"
	utilsPkg         = "github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

func findMaxCountExt(files *protoregistry.Files) (protoreflect.ExtensionType, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.ExtensionType), nil
}

func run(p *protogen.Plugin) error { _ = "STUB: not implemented"; return nil }

// Rebuild the file set so dynamic options resolve against the full graph,
// including imports we don't generate code for.

// Index every message descriptor by full name.

// A message has a Schema if it has at least one (max_count) field, or
// if it reaches a message with a Schema via a message-typed field. Find
// the closure.

// Index every protogen.Message reachable from the request by FullName.
// emit uses it to resolve a cross-file descent target's Go identifier
// from the parent generator's view of the file, rather than
// reconstructing the name from the descriptor.

// validateMaxCountValues rejects (wireguard.max_count) = 0, which would
// silently mean "no cap" at runtime (the wireguard.Scan check is
// `if rule.MaxCount > 0`). An explicit zero is almost certainly a
// mistake; pick a positive cap or drop the annotation if the field is
// genuinely unbounded.
func validateMaxCountValues(byName map[protoreflect.FullName]protoreflect.MessageDescriptor, inSchema map[protoreflect.FullName]bool, ext protoreflect.ExtensionType) error {
	_ = "STUB: not implemented"
	return nil
}

func allMDs(files *protoregistry.Files) iter.Seq[protoreflect.MessageDescriptor] {
	_ = "STUB: not implemented"
	return nil
}

func walkMsgs(mds protoreflect.MessageDescriptors) iter.Seq[protoreflect.MessageDescriptor] {
	_ = "STUB: not implemented"
	return nil
}

func hasMaxCount(d protoreflect.MessageDescriptor, ext protoreflect.ExtensionType) bool {
	_ = "STUB: not implemented"
	return false
}

func strictCheck(byName map[protoreflect.FullName]protoreflect.MessageDescriptor, inSchema map[protoreflect.FullName]bool, ext protoreflect.ExtensionType) error {
	_ = "STUB: not implemented"
	return nil
}

// emitCtx is the read-only context threaded into per-file/per-message
// emission. It batches the lookup tables and the wireguard / utils Go
// identifier expressions that emit needs but don't change between calls.
type emitCtx struct {
	// byName indexes every message descriptor (including transitive
	// imports) by FullName. Its values carry dynamic-extension options
	// resolved against the wireguard.proto descriptor, which protogen's
	// own Message.Desc does not.
	byName map[protoreflect.FullName]protoreflect.MessageDescriptor

	// byMsg indexes every protogen.Message reachable from the plugin
	// request by FullName. We use it to resolve a cross-file descent
	// target to its Go identifier without hand-constructing the name.
	byMsg map[protoreflect.FullName]*protogen.Message

	// inSchema is the set of message types we emit schemas for.
	inSchema map[protoreflect.FullName]bool

	// maxCountExt is the (wireguard.max_count) ExtensionType.
	maxCountExt protoreflect.ExtensionType
}

// emitIdents holds the Go identifier expressions for the wireguard /
// utils symbols a generated file refers to. They are produced from the
// current *protogen.GeneratedFile (which records the imports) and so are
// rebuilt for every emitted file.
type emitIdents struct {
	schema, rule, number, mustField, utilsSome string
}

func newEmitIdents(g *protogen.GeneratedFile) emitIdents {
	_ = "STUB: not implemented"
	return *new(emitIdents)
}

// emit walks files and emits per-file <name>.wireguard.go containing Schema
// vars for messages in the closure that are defined in that file.
func emit(p *protogen.Plugin, ctx emitCtx) error { _ = "STUB: not implemented"; return nil }

func emitSchema(g *protogen.GeneratedFile, m *protogen.Message, ctx emitCtx, idents emitIdents) {
	_ = "STUB: not implemented"
	// Use the descriptor from ctx.byName (which has dynamic extension
	// options resolved) rather than m.Desc (protogen's view, which
	// doesn't).
	return
}

// For a oneof variant, the wire tag is on the wrapper struct
// (e.g. Message_BlockResponse), not the parent message.

// schemaVarForTarget returns the Go expression that references the
// generated SchemaFor variable for the given message, qualified with
// the right import if it lives in a different package. We reuse the
// protogen.Message.GoIdent that the standard go generator computed
// rather than reconstructing the Go identifier from the descriptor.
func schemaVarForTarget(g *protogen.GeneratedFile, ctx emitCtx, d protoreflect.MessageDescriptor) string {
	_ = "STUB: not implemented"
	return ""
}

type pm struct{ *protogen.Message }

func (m pm) walk(yield func(pm) bool) bool { _ = "STUB: not implemented"; return false }

func allPMs(f *protogen.File) iter.Seq[pm] { _ = "STUB: not implemented"; return nil }
