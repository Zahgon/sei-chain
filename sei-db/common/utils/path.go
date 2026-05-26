package utils

// DirExists returns true if path exists and is a directory.
func DirExists(path string) bool { _ = "STUB: not implemented"; return false }

// FileExists returns true if path exists and is a regular file.
func FileExists(path string) bool { _ = "STUB: not implemented"; return false }

// GetCosmosSCStorePath returns the path for the memiavl state commitment store.
// New nodes use data/state_commit/memiavl; existing nodes with data/committer.db
// continue using the legacy path for backward compatibility.
func GetCosmosSCStorePath(homePath string) string { _ = "STUB: not implemented"; return "" }

// GetFlatKVPath returns the path for the FlatKV EVM commit store.
// New nodes use data/state_commit/flatkv; existing nodes with data/flatkv
// continue using the legacy path for backward compatibility.
func GetFlatKVPath(homePath string) string { _ = "STUB: not implemented"; return "" }

// GetStateStorePath returns the path for the Cosmos state store (SS).
// New nodes use data/state_store/cosmos/{backend}; existing nodes with
// data/{backend} continue using the legacy path for backward compatibility.
func GetStateStorePath(homePath string, backend string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetEVMStateStorePath returns the path for the EVM state store.
// New nodes use data/state_store/evm/{backend}; existing nodes with
// data/evm_ss continue using the legacy path for backward compatibility.
func GetEVMStateStorePath(homePath string, backend string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetReceiptStorePath returns the path for the receipt store.
// New nodes use data/ledger/receipt/{backend}; existing nodes with
// data/receipt.db continue using the legacy path for backward compatibility.
func GetReceiptStorePath(homePath string, backend string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetChangelogPath(dbPath string) string { _ = "STUB: not implemented"; return "" }

// ResolveAndCreateDir expands ~ to the home directory, resolves the path to
// an absolute path, and creates the directory if it doesn't exist.
func ResolveAndCreateDir(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }
