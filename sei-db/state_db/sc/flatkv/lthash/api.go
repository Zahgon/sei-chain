package lthash

import (
	"runtime"
)

// --- Public Types ---

// KVPairWithLastValue holds a KV change for LtHash computation.
type KVPairWithLastValue struct {
	Key       []byte
	Value     []byte
	LastValue []byte // Previous value (nil for new keys)
	Delete    bool   // If true, only remove last value
}

// LtHashTimings holds wall-clock timing breakdown for LtHash computation.
type LtHashTimings struct {
	TotalNs     int64
	Blake3Ns    int64
	SerializeNs int64
	MixInOutNs  int64
	MergeNs     int64
}

// DefaultLtHashWorkers defaults to NumCPU.
var DefaultLtHashWorkers = runtime.NumCPU()

// --- Public API ---

// ComputeLtHash applies changes to prev LtHash and returns the result.
// For each KV: MixOut(LastValue) if set, MixIn(Value) if not Delete.
// If prev is nil, starts from zero.
//
// Invariants consumers rely on (do NOT break these without updating
// integration tests under sei-cosmos/storev2/rootmulti that assert them):
//
//  1. Commutativity and associativity across partitions. MixIn / MixOut
//     are commutative and associative over the LtHash group, which lets
//     the parallel path below split work across N workers and merge the
//     per-worker results in any order without changing the output. Tests
//     TestFlatKVLatticeHashDeterminism and
//     TestFlatKVLargeChangesetDeterminism depend on this.
//
//  2. Delete-of-absent-key is a no-op. When LastValue is nil (key was not
//     previously present) and Delete is true, both lastSerialized and
//     newSerialized remain nil, so neither MixOut nor MixIn is invoked and
//     this entry contributes zero to the hash. Same-block set-then-delete
//     of a non-existent key therefore cannot shift the LtHash.
//     TestFlatKVDeleteAndOverwriteWorkload (block 5) depends on this.
func ComputeLtHash(prev *LtHash, kvPairs []KVPairWithLastValue) (*LtHash, *LtHashTimings) {
	_ = "STUB: not implemented"
	return nil, nil
}

// --- Internal computation ---

// serializedKV holds serialized key-value data for hashing.
type serializedKV struct {
	lastSerialized []byte
	newSerialized  []byte
}

// lthashPair holds computed LtHash values for a single KV change.
type lthashPair struct {
	lastLth *LtHash
	newLth  *LtHash
}

// computeDelta computes the LtHash delta for a changeset.
func computeDelta(kvPairs []KVPairWithLastValue, numWorkers int) (*LtHash, *LtHashTimings) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Small changesets: serial is faster

// Phase 1: Serialize

// Phase 2: Hash (parallel)

// Phase 3: MixIn/MixOut (parallel)

// Phase 4: Merge

// computeDeltaSerial is the serial version for small changesets.
func computeDeltaSerial(kvPairs []KVPairWithLastValue) (*LtHash, *LtHashTimings) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Phase 1: Serialize

// Phase 2: Hash

// Phase 3: MixIn/MixOut
