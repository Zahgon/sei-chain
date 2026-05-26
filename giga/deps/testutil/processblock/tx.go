package processblock

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	xauthsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/tx"
)

var InterfaceReg = types.NewInterfaceRegistry()
var Marshaler = codec.NewProtoCodec(InterfaceReg)
var TxConfig = tx.NewTxConfig(Marshaler, tx.DefaultSignModes)

func (a *App) Sign(account sdk.AccAddress, fee int64, msgs ...sdk.Msg) xauthsigning.Tx {
	_ = "STUB: not implemented"
	return *new(xauthsigning.Tx)
}

func GetKey(mnemonic string) cryptotypes.PrivKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PrivKey)
}
