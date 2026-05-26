package flatkv

import (
	seidbtypes "github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/vtype"
)

// Get returns the value for the given key within the specified module.
// For EVM keys (moduleName == keys.EVMStoreKey), the key is a prefix-encoded
// EVM key routed internally to account/storage/code/legacy DBs.
// For non-EVM modules, the key is read from legacy storage with the module prefix.
// Returns (value, true) if found, (nil, false) if not found.
// Panics on I/O errors or unsupported key types.
func (s *CommitStore) Get(moduleName string, key []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// CodeHash

// GetBlockHeightModified returns the block height at which the key was last modified.
// Only supported for EVM keys; non-EVM legacy data does not track block height.
// If not found, returns (-1, false, nil).
func (s *CommitStore) GetBlockHeightModified(moduleName string, key []byte) (int64, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// Has reports whether the key exists within the given module.
// Panics on I/O errors or unsupported key types.
func (s *CommitStore) Has(moduleName string, key []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// =============================================================================
// Internal Getters (used by ApplyChangeSets for LtHash computation)
// =============================================================================

// readFromDB checks pending writes first, then falls back to a DB read.
// Returns (zero, nil) when the key is not found.
func readFromDB[T vtype.VType](
	physKey []byte,
	pendingWrites map[string]T,
	db seidbtypes.KeyValueDB,
	deserialize func([]byte) (T, error),
	dbName string,
) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (s *CommitStore) getAccountData(keyBytes []byte) (*vtype.AccountData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CommitStore) getStorageData(keyBytes []byte) (*vtype.StorageData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CommitStore) getStorageValue(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CommitStore) getCodeData(keyBytes []byte) (*vtype.CodeData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CommitStore) getCodeValue(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CommitStore) getLegacyData(moduleName string, keyBytes []byte) (*vtype.LegacyData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CommitStore) getLegacyValue(moduleName string, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
