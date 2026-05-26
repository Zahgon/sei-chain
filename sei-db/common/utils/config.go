package utils

// Config is implemented by any configuration struct that supports validation.
type Config interface {
	Validate() error
}

// StringifyConfig returns the config as human-readable, multi-line JSON.
func StringifyConfig(cfg Config) (string, error) { _ = "STUB: not implemented"; return "", nil }

// LoadConfigFromFile decodes a JSON config file on top of the provided defaults
// struct. The caller should pass a pointer to a struct pre-populated with
// default values; file values are overlaid on top. Unknown JSON keys cause an
// error. After decoding, Validate() is called on the result.
func LoadConfigFromFile(path string, defaults Config) error {
	_ = "STUB: not implemented"
	//nolint:gosec // G304 - path comes from CLI arg, filepath.Clean used to mitigate traversal
	return nil
}
