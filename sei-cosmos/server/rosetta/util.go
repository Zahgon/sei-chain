package rosetta

import (
	"time"
)

// timeToMilliseconds converts time to milliseconds timestamp
func timeToMilliseconds(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

// unmarshalMetadata unmarshals the given meta to the target
func unmarshalMetadata(meta map[string]interface{}, target interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// marshalMetadata marshals the given interface to map[string]interface{}
func marshalMetadata(o interface{}) (meta map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
