// Package keymigrate translates all legacy formatted keys to their
// new components.
//
// The key migration operation as implemented provides a potential
// model for database migration operations. Crucially, the migration
// as implemented does not depend on any tendermint code.
package keymigrate

import (
	"context"

	dbm "github.com/tendermint/tm-db"
)

type (
	keyID       []byte
	migrateFunc func(keyID) (keyID, error)
)

func getAllLegacyKeys(db dbm.DB) ([]keyID, error) { _ = "STUB: not implemented"; return nil, nil }

// make sure it's a key with a legacy format, and skip
// all other keys, to make it safe to resume the migration.

// Make an explicit copy, since not all tm-db backends do.

// keyType is an enumeration for the structural type of a key.
type keyType int

func (t keyType) isLegacy() bool { _ = "STUB: not implemented"; return false }

const (
	nonLegacyKey keyType = iota // non-legacy key (presumed already converted)
	consensusParamsKey
	abciResponsesKey
	validatorsKey
	stateStoreKey        // state storage record
	blockMetaKey         // H:
	blockPartKey         // P:
	commitKey            // C:
	seenCommitKey        // SC:
	blockHashKey         // BH:
	lightSizeKey         // size
	lightBlockKey        // lb/
	evidenceCommittedKey // \x00
	evidencePendingKey   // \x01
	txHeightKey          // tx.height/... (special case)
	abciEventKey         // name/value/height/index
	txHashKey            // 32-byte transaction hash (unprefixed)
)

var prefixes = []struct {
	prefix []byte
	ktype  keyType
	check  func(keyID) bool
}{
	{[]byte("consensusParamsKey:"), consensusParamsKey, nil},
	{[]byte("abciResponsesKey:"), abciResponsesKey, nil},
	{[]byte("validatorsKey:"), validatorsKey, nil},
	{[]byte("stateKey"), stateStoreKey, nil},
	{[]byte("H:"), blockMetaKey, nil},
	{[]byte("P:"), blockPartKey, nil},
	{[]byte("C:"), commitKey, nil},
	{[]byte("SC:"), seenCommitKey, nil},
	{[]byte("BH:"), blockHashKey, nil},
	{[]byte("size"), lightSizeKey, nil},
	{[]byte("lb/"), lightBlockKey, nil},
	{[]byte("\x00"), evidenceCommittedKey, checkEvidenceKey},
	{[]byte("\x01"), evidencePendingKey, checkEvidenceKey},
}

// checkKeyType classifies a candidate key based on its structure.
func checkKeyType(key keyID) keyType { _ = "STUB: not implemented"; return *new(keyType) }

// A legacy event key has the form:
//
//    <name> / <value> / <height> / <index>
//
// Transaction hashes are stored as a raw binary hash with no prefix.
//
// Because a hash can contain any byte, it is possible (though unlikely)
// that a hash could have the correct form for an event key, in which case
// we would translate it incorrectly.  To reduce the likelihood of an
// incorrect interpretation, we parse candidate event keys and check for
// some structural properties before making a decision.
//
// Note, though, that nothing prevents event names or values from containing
// additional "/" separators, so the parse has to be forgiving.

// Special case for tx.height.

// The name cannot be empty, but we don't know where the name ends and
// the value begins, so insist that there be something.

// Check whether the last two fields could be .../height/index.

// If we get here, it's not an event key. Treat it as a hash if it is the
// right length. Note that it IS possible this could collide with the
// translation of some other key (though not a hash, since encoded hashes
// will be longer). The chance of that is small, but there is nothing we can
// do to detect it.

// isDecimal reports whether buf is a non-empty sequence of Unicode decimal
// digits.
func isDecimal(buf []byte) bool { _ = "STUB: not implemented"; return false }

func migrateKey(key keyID) (keyID, error) { _ = "STUB: not implemented"; return *new(keyID), nil }

// drop prefix

func convertEvidence(key keyID, newPrefix int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkEvidenceKey reports whether a candidate key with one of the legacy
// evidence prefixes has the correct structure for a legacy evidence key.
//
// This check is needed because transaction hashes are stored without a prefix,
// so checking the one-byte prefix alone is not enough to distinguish them.
// Legacy evidence keys are suffixed with a string of the format:
//
//	"%0.16X/%X"
//
// where the first element is the height and the second is the hash.  Thus, we
// check
func checkEvidenceKey(key keyID) bool { _ = "STUB: not implemented"; return false }

func isHex(data []byte) bool { _ = "STUB: not implemented"; return false }

func replaceKey(db dbm.DB, key keyID, gooseFn migrateFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// 10% of the time, force a write to disk, but mostly don't,
// because it's faster.
// nolint:gosec

// Migrate converts all legacy key formats to new key formats. The
// operation is idempotent, so it's safe to resume a failed
// operation. The operation is somewhat parallelized, relying on the
// concurrency safety of the underlying databases.
//
// Migrate has "continue on error" semantics and will iterate through
// all legacy keys attempt to migrate them, and will collect all
// errors and will return only at the end of the operation.
//
// The context allows for a safe termination of the operation
// (e.g connected to a singal handler,) to abort the operation
// in-between migration operations.
func Migrate(ctx context.Context, db dbm.DB) error { _ = "STUB: not implemented"; return nil }
