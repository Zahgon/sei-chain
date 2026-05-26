package types

import (
	"reflect"

	"github.com/gogo/protobuf/proto"
)

type legacyInterfaceRegistry struct {
	interfaceNames map[string]reflect.Type
	interfaceImpls map[reflect.Type]interfaceMap
	typeURLMap     map[string]reflect.Type
}

// NewInterfaceRegistry returns a new InterfaceRegistry
func NewLegacyInterfaceRegistry() InterfaceRegistry {
	_ = "STUB: not implemented"
	return *new(InterfaceRegistry)
}

func (registry *legacyInterfaceRegistry) RegisterInterface(protoName string, iface interface{}, impls ...proto.Message) {
	_ = "STUB: not implemented"
	return
}

// RegisterImplementations registers a concrete proto Message which implements
// the given interface.
//
// This function PANICs if different concrete types are registered under the
// same typeURL.
func (registry *legacyInterfaceRegistry) RegisterImplementations(iface interface{}, impls ...proto.Message) {
	_ = "STUB: not implemented"
	return
}

// RegisterCustomTypeURL registers a concrete type which implements the given
// interface under `typeURL`.
//
// This function PANICs if different concrete types are registered under the
// same typeURL.
func (registry *legacyInterfaceRegistry) RegisterCustomTypeURL(iface interface{}, typeURL string, impl proto.Message) {
	_ = "STUB: not implemented"
	return
}

// registerImpl registers a concrete type which implements the given
// interface under `typeURL`.
//
// This function PANICs if different concrete types are registered under the
// same typeURL.
func (registry *legacyInterfaceRegistry) registerImpl(iface interface{}, typeURL string, impl proto.Message) {
	_ = "STUB: not implemented"
	return
}

// Check if we already registered something under the given typeURL. It's
// okay to register the same concrete type again, but if we are registering
// a new concrete type under the same typeURL, then we throw an error (here,
// we panic).

func (registry *legacyInterfaceRegistry) ListAllInterfaces() []string {
	_ = "STUB: not implemented"
	return nil
}

func (registry *legacyInterfaceRegistry) ListImplementations(ifaceName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (registry *legacyInterfaceRegistry) UnpackAny(any *Any, iface interface{}) error {
	_ = "STUB: not implemented"
	// here we gracefully handle the case in which `any` itself is `nil`, which may occur in message decoding
	return nil
}

// if TypeUrl is empty return nil because without it we can't actually unpack anything

// Resolve returns the proto message given its typeURL. It works with types
// registered with RegisterInterface/RegisterImplementations, as well as those
// registered with RegisterWithCustomTypeURL.
func (registry *legacyInterfaceRegistry) Resolve(typeURL string) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}
