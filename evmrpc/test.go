package evmrpc

type TestAPI struct{}

func NewTestAPI() *TestAPI { _ = "STUB: not implemented"; return nil }

func (a *TestAPI) IncrementPointerVersion(pointerType string, offset int16) error {
	_ = "STUB: not implemented"
	return nil
}
