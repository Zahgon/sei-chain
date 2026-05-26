package flatkv

import (
	"github.com/sei-protocol/sei-chain/sei-db/common/keys"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/lthash"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/vtype"
)

// ApplyChangeSets buffers EVM changesets and updates LtHash.
// Non-EVM modules are routed to legacyDB with a "<module>/" key prefix.
func (s *CommitStore) ApplyChangeSets(changeSets []*proto.NamedChangeSet) (err error) {
	_ = "STUB: not implemented"
	return nil
}

///////////
// Setup //
///////////

////////////////////
// Batch Read Old //
////////////////////

//////////////////
// Gather Pairs //
//////////////////

// Gather account pairs

// TODO: update this when we add a balance key!

////////////////////
// Compute LTHash //
////////////////////

// Global LTHash = sum of per-DB hashes (homomorphic property).
// Compute into a fresh hash and swap to avoid a transient empty state
// on workingLtHash (safe for future pipelining / async callers).

//////////////
// Finalize //
//////////////

// Now that we've made it through the batch without errors, we can add the change sets to the pending change sets.

// classifyAndPrefix splits changeSets into per-EVMKeyKind maps whose keys are
// already in physical format ("module/" + prefix_encoded_key). Non-EVM modules are
// merged into the EVMKeyLegacy bucket with a "<module>/" prefix.
//
// This replaces the former sortChangeSets + prefixModuleKeys two-pass approach,
// avoiding an extra map allocation and repeated string concatenation per key.
func classifyAndPrefix(changeSets []*proto.NamedChangeSet) (map[keys.EVMKeyKind]map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process incoming storage changes into a form appropriate for hashing and insertion into the DB.
func processStorageChanges(
	rawChanges map[string][]byte,
	blockHeight int64,
) (map[string]*vtype.StorageData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletion is equivalent to setting the storage value to a zero value

// Process incoming code changes into a form appropriate for hashing and insertion into the DB.
func processCodeChanges(
	rawChanges map[string][]byte,
	blockHeight int64,
) (map[string]*vtype.CodeData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletion is equivalent to setting the code to a zero value

// Process incoming legacy changes into a form appropriate for hashing and insertion into the DB.
func processLegacyChanges(
	rawChanges map[string][]byte,
	blockHeight int64,
) (map[string]*vtype.LegacyData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gatherLTHashPairs[T vtype.VType](
	newValues map[string]T,
	oldValues map[string]T,
) []lthash.KVPairWithLastValue {
	_ = "STUB: not implemented"
	return nil
}

// Merge account updates down into a single update per account.
func mergeAccountUpdates(
	nonceChanges map[string][]byte,
	codeHashChanges map[string][]byte,
	balanceChanges map[string][]byte,
) (map[string]*vtype.PendingAccountWrite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deletion is equivalent to setting the nonce to 0

// Deletion is equivalent to setting the code hash to a zero hash

// Deletion is equivalent to setting the balance to a zero balance

// Combine the pending account writes with prior values to determine the new account values.
//
// We need to take this step because accounts are split into multiple fields, and it's possible to overwrite just a
// single field (thus requiring us to copy the unmodified fields from the prior value).
func deriveNewAccountValues(
	pendingWrites map[string]*vtype.PendingAccountWrite,
	oldValues map[string]*vtype.AccountData,
	blockHeight int64,
) map[string]*vtype.AccountData {
	_ = "STUB: not implemented"
	return nil
}
