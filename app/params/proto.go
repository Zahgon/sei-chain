//go:build !test_amino

package params

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() EncodingConfig { _ = "STUB: not implemented"; return *new(EncodingConfig) }

// MakeLegacyEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeLegacyEncodingConfig() EncodingConfig {
	_ = "STUB: not implemented"
	return *new(EncodingConfig)
}
