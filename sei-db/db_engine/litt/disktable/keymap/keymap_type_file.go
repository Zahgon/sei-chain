package keymap

// KeymapTypeFileName is the name of the file that contains the keymap type.
const KeymapTypeFileName = "keymap-type.txt"

// KeymapTypeFile is a text file that contains the name of the keymap type. This is used to determine if the keymap
// needs to reload when littDB is restarted, or if the data structures in the keymap directory are still valid.
type KeymapTypeFile struct {
	// keymapPath is the path to the keymap directory.
	keymapPath string

	// KeymapType is the type of the keymap currently stored in the keymap directory.
	keymapType KeymapType
}

// KeymapFileExists checks if the keymap type file exists in the target directory.
func KeymapFileExists(keymapPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NewKeymapTypeFile creates a new KeymapTypeFile.
func NewKeymapTypeFile(keymapPath string, keymapType KeymapType) *KeymapTypeFile {
	_ = "STUB: not implemented"
	return nil
}

// LoadKeymapTypeFile loads the keymap type from the keymap directory.
func LoadKeymapTypeFile(keymapPath string) (*KeymapTypeFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // path within keymap directory

// Type returns the type of the keymap.
func (k *KeymapTypeFile) Type() KeymapType {
	_ = "STUB: not implemented"
	return *

	// Write writes the keymap type to the keymap directory.
	new(KeymapType)
}

func (k *KeymapTypeFile) Write() error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // path within keymap directory

// Delete deletes the keymap type file.
func (k *KeymapTypeFile) Delete() error { _ = "STUB: not implemented"; return nil }
