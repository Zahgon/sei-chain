package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/slashing/types"
)

// Keeper of the slashing store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	sk         types.StakingKeeper
	paramspace types.ParamSubspace
}

// NewKeeper creates a slashing keeper
func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, sk types.StakingKeeper, paramspace types.ParamSubspace) Keeper {
	_ = "STUB: not implemented"
	// set KeyTable if it has not already been set
	return *new(Keeper)
}

// AddPubkey sets a address-pubkey relation
func (k Keeper) AddPubkey(ctx sdk.Context, pubkey cryptotypes.PubKey) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPubkey returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetPubkey(ctx sdk.Context, a cryptotypes.Address) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// Slash attempts to slash a validator. The slash is delegated to the staking
// module to make the necessary validator changes.
func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, fraction sdk.Dec, power, distributionHeight int64) {
	_ = "STUB: not implemented"
	return
}

// Jail attempts to jail a validator. The slash is delegated to the staking module
// to make the necessary validator changes.
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) { _ = "STUB: not implemented"; return }

func (k Keeper) deleteAddrPubkeyRelation(ctx sdk.Context, addr cryptotypes.Address) {
	_ = "STUB: not implemented"
	return
}

func (k Keeper) GetStoreKey() sdk.StoreKey { _ = "STUB: not implemented"; return *new(sdk.StoreKey) }
