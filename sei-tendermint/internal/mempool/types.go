package mempool

const (
	// UnknownPeerID is the peer ID to use when running CheckTx when there is
	// no peer (e.g. RPC)
	UnknownPeerID uint16 = 0
)

// TxConstraints contains the precomputed consensus-derived mempool limits for
// the current state snapshot.
type TxConstraints struct {
	MaxDataBytes int64
	MaxGas       int64
}

// TxConstraintsFetcher returns the precomputed consensus-derived mempool limits for the current
// state snapshot.
type TxConstraintsFetcher func() (TxConstraints, error)

func NopTxConstraintsFetcher() (TxConstraints, error) {
	_ = "STUB: not implemented"
	return *new(TxConstraints), nil
}
