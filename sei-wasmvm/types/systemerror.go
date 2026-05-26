package types

// SystemError captures all errors returned from the Rust code as SystemError.
// Exactly one of the fields should be set.
type SystemError struct {
	InvalidRequest     *InvalidRequest     `json:"invalid_request,omitempty"`
	InvalidResponse    *InvalidResponse    `json:"invalid_response,omitempty"`
	NoSuchContract     *NoSuchContract     `json:"no_such_contract,omitempty"`
	NoSuchCode         *NoSuchCode         `json:"no_such_code,omitempty"`
	Unknown            *Unknown            `json:"unknown,omitempty"`
	UnsupportedRequest *UnsupportedRequest `json:"unsupported_request,omitempty"`
}

var (
	_ error = SystemError{}
	_ error = InvalidRequest{}
	_ error = InvalidResponse{}
	_ error = NoSuchContract{}
	_ error = Unknown{}
	_ error = UnsupportedRequest{}
)

func (a SystemError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidRequest struct {
	Err     string `json:"error"`
	Request []byte `json:"request"`
}

func (e InvalidRequest) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidResponse struct {
	Err      string `json:"error"`
	Response []byte `json:"response"`
}

func (e InvalidResponse) Error() string { _ = "STUB: not implemented"; return "" }

type NoSuchContract struct {
	Addr string `json:"addr,omitempty"`
}

func (e NoSuchContract) Error() string { _ = "STUB: not implemented"; return "" }

type NoSuchCode struct {
	CodeID uint64 `json:"code_id,omitempty"`
}

func (e NoSuchCode) Error() string { _ = "STUB: not implemented"; return "" }

type Unknown struct{}

func (e Unknown) Error() string { _ = "STUB: not implemented"; return "" }

type UnsupportedRequest struct {
	Kind string `json:"kind,omitempty"`
}

func (e UnsupportedRequest) Error() string { _ = "STUB: not implemented"; return "" }

// ToSystemError will try to convert the given error to an SystemError.
// This is important to returning any Go error back to Rust.
//
// If it is already StdError, return self.
// If it is an error, which could be a sub-field of StdError, embed it.
// If it is anything else, **return nil**
//
// This may return nil on an unknown error, whereas ToStdError will always create
// a valid error type.
func ToSystemError(err error) *SystemError { _ = "STUB: not implemented"; return nil }

// check if an interface is nil (even if it has type info)
func isNil(i interface{}) bool { _ = "STUB: not implemented"; return false }

// IsNil panics if you try it on a struct (not a pointer)

// if we aren't a pointer, can't be nil, can we?
