package codec

import (
	"github.com/gogo/protobuf/proto"
)

// MarshalYAML marshals toPrint using JSONCodec to leverage specialized MarshalJSON methods
// (usually related to serialize data with protobuf or amin depending on a configuration).
// This involves additional roundtrip through JSON.
func MarshalYAML(cdc JSONCodec, toPrint proto.Message) ([]byte, error) {
	_ = "STUB: not implemented"
	// We are OK with the performance hit of the additional JSON roundtip. MarshalYAML is not
	// used in any critical parts of the system.
	return nil, nil
}

// generate YAML by decoding JSON and re-encoding to YAML
