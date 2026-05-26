package types

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto/merkle"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
)

// Tx is an arbitrary byte array.
// NOTE: Tx has no types at this level, so when wire encoded it's just length-prefixed.
// Might we want types here ?
type Tx []byte

// Hash computes the TMHASH hash of the wire encoded transaction.
func (tx Tx) Hash() TxHash { _ = "STUB: not implemented"; return *new(TxHash) }

// String returns the hex-encoded transaction as a string.
func (tx Tx) String() string { _ = "STUB: not implemented"; return "" }

// Txs is a slice of Tx.
type Txs []Tx

// Hash returns the Merkle root hash of the transaction hashes.
// i.e. the leaves of the tree are the hashes of the txs.
func (txs Txs) Hash() []byte { _ = "STUB: not implemented"; return nil }

// Index returns the index of this transaction in the list, or -1 if not found
func (txs Txs) Index(tx Tx) int { _ = "STUB: not implemented"; return 0 }

// IndexByHash returns the index of this transaction hash in the list, or -1 if not found
func (txs Txs) IndexByHash(hash TxHash) int { _ = "STUB: not implemented"; return 0 }

func (txs Txs) Proof(i int) TxProof { _ = "STUB: not implemented"; return *new(TxProof) }

func (txs Txs) hashList() [][]byte { _ = "STUB: not implemented"; return nil }

// Txs is a slice of transactions. Sorting a Txs value orders the transactions
// lexicographically.
func (txs Txs) Len() int           { _ = "STUB: not implemented"; return 0 }
func (txs Txs) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (txs Txs) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// ToSliceOfBytes converts a Txs to slice of byte slices.
func (txs Txs) ToSliceOfBytes() [][]byte { _ = "STUB: not implemented"; return nil }

// TxProof represents a Merkle proof of the presence of a transaction in the Merkle tree.
type TxProof struct {
	RootHash tmbytes.HexBytes `json:"root_hash"`
	Data     Tx               `json:"data"`
	Proof    merkle.Proof     `json:"proof"`
}

// Leaf returns the hash(tx), which is the leaf in the merkle tree which this proof refers to.
func (tp TxProof) Leaf() []byte { _ = "STUB: not implemented"; return nil }

// Validate verifies the proof. It returns nil if the RootHash matches the dataHash argument,
// and if the proof is internally consistent. Otherwise, it returns a sensible error.
func (tp TxProof) Validate(dataHash []byte) error { _ = "STUB: not implemented"; return nil }

func (tp TxProof) ToProto() tmproto.TxProof {
	_ = "STUB: not implemented"
	return *new(tmproto.TxProof)
}

func TxProofFromProto(pb tmproto.TxProof) (TxProof, error) {
	_ = "STUB: not implemented"
	return *new(TxProof), nil
}

// ComputeProtoSizeForTxs wraps the transactions in tmproto.Data{} and calculates the size.
// https://developers.google.com/protocol-buffers/docs/encoding
func ComputeProtoSizeForTxs(txs []Tx) int64 { _ = "STUB: not implemented"; return 0 }
