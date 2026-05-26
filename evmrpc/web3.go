package evmrpc

type Web3API struct{}

func (w *Web3API) ClientVersion() string {
	_ = "STUB: not implemented"
	// Sei EVM is backed by go-ethereum
	return ""
}
