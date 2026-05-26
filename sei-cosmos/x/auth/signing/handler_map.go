package signing

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// SignModeHandlerMap is SignModeHandler that aggregates multiple SignModeHandler's into
// a single handler
type SignModeHandlerMap struct {
	defaultMode      signing.SignMode
	modes            []signing.SignMode
	signModeHandlers map[signing.SignMode]SignModeHandler
}

var _ SignModeHandler = SignModeHandlerMap{}

// NewSignModeHandlerMap returns a new SignModeHandlerMap with the provided defaultMode and handlers
func NewSignModeHandlerMap(defaultMode signing.SignMode, handlers []SignModeHandler) SignModeHandlerMap {
	_ = "STUB: not implemented"
	return *new(SignModeHandlerMap)
}

//nolint:prealloc // Not worth pre-allocation, considering the nested loop

// DefaultMode implements SignModeHandler.DefaultMode
func (h SignModeHandlerMap) DefaultMode() signing.SignMode {
	_ = "STUB: not implemented"
	return *

	// Modes implements SignModeHandler.Modes
	new(signing.SignMode)
}

func (h SignModeHandlerMap) Modes() []signing.SignMode {
	_ = "STUB: not implemented"

	// DefaultMode implements SignModeHandler.GetSignBytes
	return nil
}

func (h SignModeHandlerMap) GetSignBytes(mode signing.SignMode, data SignerData, tx sdk.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
