package state

type State struct {
	LastProcessedHeight int64   `json:"last_processed_height"`
	BlocksMissingTxs    []int64 `json:"blocks_missing_txs"`
}

// WriteState write the state to a JSON file.
func WriteState(dir string, s State) error { _ = "STUB: not implemented"; return nil }

// ReadState reads the state from a JSON file.
func ReadState(dir string) (State, error) { _ = "STUB: not implemented"; return *new(State), nil }
