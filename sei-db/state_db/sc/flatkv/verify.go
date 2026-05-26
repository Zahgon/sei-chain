package flatkv

// VerifyLtHash full-scans all four data DBs and checks the recomputed LtHash
// against the store's committedLtHash. Read-write stores with uncommitted
// ApplyChangeSets writes are rejected (the on-disk scan cannot see them).
//
// Buffers every KV in memory (peak RSS ~2-3x on-disk size) and is not
// cancellable. Intended for tests and offline maintenance / cutover checks;
// not suitable for online verification of production-sized state.
func VerifyLtHash(s Store) error { _ = "STUB: not implemented"; return nil }

func verifyLtHashInternal(cs *CommitStore) error {
	_ = "STUB: not implemented"
	// A read-write store between ApplyChangeSets and Commit has
	// workingLtHash != committedLtHash. The full scan below reads only
	// persisted DB contents, so there is no way to validate the in-memory
	// pending state against disk here. Fail loudly rather than masquerade
	// a pending-writes situation as an integrity error.
	return nil
}

// Full scan reflects on-disk (committed) state, so the only correct
// reference is committedLtHash. workingLtHash may include uncommitted
// ApplyChangeSets updates that have not yet been persisted.
