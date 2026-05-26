package evmrpc

type EchoAPI struct{}

func NewEchoAPI() *EchoAPI { _ = "STUB: not implemented"; return nil }

func (a *EchoAPI) Echo(data string) string { _ = "STUB: not implemented"; return "" }
