package crisis

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/crisis/types"
)

// RouterKey
const RouterKey = types.ModuleName

func NewHandler(k types.MsgServer) sdk.Handler { _ = "STUB: not implemented"; return *new(sdk.Handler) }
