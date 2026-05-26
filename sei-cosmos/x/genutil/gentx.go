package genutil

import (
	"encoding/json"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
)

// SetGenTxsInAppGenesisState - sets the genesis transactions in the app genesis state
func SetGenTxsInAppGenesisState(
	cdc codec.JSONCodec, txJSONEncoder sdk.TxEncoder, appGenesisState map[string]json.RawMessage, genTxs []sdk.Tx,
) (map[string]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ValidateAccountInGenesis checks that the provided account has a sufficient
// balance in the set of genesis accounts.
func ValidateAccountInGenesis(
	appGenesisState map[string]json.RawMessage, genBalIterator types.GenesisBalancesIterator,
	addr sdk.Address, coins sdk.Coins, cdc codec.JSONCodec,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ensure that account is in genesis

// ensure account contains enough funds of default bond denom

type deliverTxfn func(sdk.Context, abci.RequestDeliverTxV2, sdk.Tx, [32]byte) abci.ResponseDeliverTx

// DeliverGenTxs iterates over all genesis txs, decodes each into a Tx and
// invokes the provided deliverTxfn with the decoded Tx. It returns the result
// of the staking module's ApplyAndReturnValidatorSetUpdates.
func DeliverGenTxs(
	ctx sdk.Context, genTxs []json.RawMessage,
	stakingKeeper types.StakingKeeper, deliverTx deliverTxfn,
	txEncodingConfig client.TxEncodingConfig,
) ([]abci.ValidatorUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
