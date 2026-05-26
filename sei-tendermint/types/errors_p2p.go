package types

//-------------------------------------------------------------------

type ErrNetAddressNoID struct {
	Addr string
}

func (e ErrNetAddressNoID) Error() string { _ = "STUB: not implemented"; return "" }

type ErrNetAddressInvalid struct {
	Addr string
	Err  error
}

func (e ErrNetAddressInvalid) Error() string { _ = "STUB: not implemented"; return "" }

type ErrNetAddressLookup struct {
	Addr string
	Err  error
}

func (e ErrNetAddressLookup) Error() string { _ = "STUB: not implemented"; return "" }
