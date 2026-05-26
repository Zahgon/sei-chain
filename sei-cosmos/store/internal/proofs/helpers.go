package proofs

import (
	tmcrypto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
)

// SimpleResult contains a merkle.SimpleProof along with all data needed to build the confio/proof
type SimpleResult struct {
	Key      []byte
	Value    []byte
	Proof    *tmcrypto.Proof
	RootHash []byte
}

// GenerateRangeProof makes a tree of size and returns a range proof for one random element
//
// returns a range proof and the root hash of the tree
func GenerateRangeProof(size int, loc Where) *SimpleResult { _ = "STUB: not implemented"; return nil }

// Where selects a location for a key - Left, Right, or Middle
type Where int

const (
	Left Where = iota
	Right
	Middle
)

func SortedKeys(data map[string][]byte) []string { _ = "STUB: not implemented"; return nil }

func CalcRoot(data map[string][]byte) []byte { _ = "STUB: not implemented"; return nil }

// GetKey this returns a key, on Left/Right/Middle
func GetKey(allkeys []string, loc Where) string { _ = "STUB: not implemented"; return "" }

// select a random index between 1 and allkeys-2

// GetNonKey returns a missing key - Left of all, Right of all, or in the Middle
func GetNonKey(allkeys []string, loc Where) string { _ = "STUB: not implemented"; return "" }

// otherwise, next to an existing key (copy before mod)

func toValue(key string) []byte { _ = "STUB: not implemented"; return nil }

// BuildMap creates random key/values and stores in a map,
// returns a list of all keys in sorted order
func BuildMap(size int) map[string][]byte { _ = "STUB: not implemented"; return nil }

// insert lots of info and store the bytes
