package genutil

// DONTCOVER

import (
	"encoding/json"

	cfg "github.com/sei-protocol/sei-chain/sei-tendermint/config"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/genutil/types"
)

// GenAppStateFromConfig gets the genesis app state from the config
func GenAppStateFromConfig(cdc codec.JSONCodec, txEncodingConfig client.TxEncodingConfig,
	config *cfg.Config, initCfg types.InitConfig, genDoc tmtypes.GenesisDoc, genBalIterator types.GenesisBalancesIterator,
) (appState json.RawMessage, err error) {
	_ = "STUB: not implemented"

	// process genesis transactions, else create default genesis.json
	return *new(json.RawMessage), nil
}

// if there are no gen txs to be processed, return the default empty state

// create the app state

// CollectTxs processes and validates application's genesis Txs and returns
// the list of appGenTxs, and persistent peers required to generate genesis.json.
func CollectTxs(cdc codec.JSONCodec, txJSONDecoder sdk.TxDecoder, moniker, genTxsDir string,
	genDoc tmtypes.GenesisDoc, genBalIterator types.GenesisBalancesIterator,
) (appGenTxs []sdk.Tx, persistentPeers string, err error) {
	_ = "STUB: not implemented"
	// prepare a map of all balances in genesis state to then validate
	// against the validators addresses
	return nil, "", nil
}

// addresses and IPs (and port) validator server info

// get the genTx

// the memo flag is used to store
// the ip and node-id, for example this may be:
// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"

// genesis transactions must be single-message

// TODO abstract out staking message validation back to staking

// validate delegator and validator addresses and funds against the accounts in the state

// exclude itself from persistent peers
