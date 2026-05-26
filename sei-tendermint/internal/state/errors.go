package state

type (
	ErrInvalidBlock error
	ErrProxyAppConn error

	ErrUnknownBlock struct {
		Height int64
	}

	ErrBlockHashMismatch struct {
		CoreHash []byte
		AppHash  []byte
		Height   int64
	}

	ErrAppBlockHeightTooHigh struct {
		CoreHeight int64
		AppHeight  int64
	}

	ErrAppBlockHeightTooLow struct {
		AppHeight int64
		StoreBase int64
	}

	ErrLastStateMismatch struct {
		Height int64
		Core   []byte
		App    []byte
	}

	ErrStateMismatch struct {
		Got      *State
		Expected *State
	}

	ErrNoValSetForHeight struct {
		Height int64
		Err    error
	}

	ErrNoConsensusParamsForHeight struct {
		Height int64
	}

	ErrNoFinalizeBlockResponsesForHeight struct {
		Height int64
	}
)

func (e ErrUnknownBlock) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrBlockHashMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrAppBlockHeightTooHigh) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrAppBlockHeightTooLow) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrLastStateMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrStateMismatch) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrNoValSetForHeight) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrNoValSetForHeight) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e ErrNoConsensusParamsForHeight) Error() string { _ = "STUB: not implemented"; return "" }

func (e ErrNoFinalizeBlockResponsesForHeight) Error() string { _ = "STUB: not implemented"; return "" }
