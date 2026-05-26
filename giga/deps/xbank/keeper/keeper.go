package keeper

import (
	"github.com/sei-protocol/sei-chain/giga/deps/xbank/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/seilog"
)

var (
	logger = seilog.NewLogger("giga", "deps", "xbank", "keeper")

	_ Keeper = (*BaseKeeper)(nil)
)

// Keeper defines a module interface that facilitates the transfer of coins
// between accounts.
type Keeper interface {
	SendKeeper

	InitGenesis(sdk.Context, *types.GenesisState)
	ExportGenesis(sdk.Context) *types.GenesisState

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	HasSupply(ctx sdk.Context, denom string) bool
	SetSupply(ctx sdk.Context, coin sdk.Coin)
	GetDenomMetaData(ctx sdk.Context, denom string) (types.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData types.Metadata)

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error

	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	DeferredSendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	WriteDeferredBalances(ctx sdk.Context) []abci.Event
	IterateDeferredBalances(ctx sdk.Context, cb func(addr sdk.AccAddress, coin sdk.Coin) bool)

	DelegateCoins(ctx sdk.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoins(ctx sdk.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error

	GetStoreKey() sdk.StoreKey
	GetCdc() codec.BinaryCodec
}

// BaseKeeper manages transfers between accounts. It implements the Keeper interface.
type BaseKeeper struct {
	BaseSendKeeper

	ak                     types.AccountKeeper
	deferredCache          *DeferredCache
	cdc                    codec.BinaryCodec
	storeKey               sdk.StoreKey
	paramSpace             paramtypes.Subspace
	mintCoinsRestrictionFn MintingRestrictionFn
	cacheSize              int
}

type MintingRestrictionFn func(ctx sdk.Context, coins sdk.Coins) error
type SubFn func(ctx sdk.Context, moduleName string, amounts sdk.Coins) error
type AddFn func(ctx sdk.Context, moduleName string, amounts sdk.Coins) error

// NewBaseKeeper returns a new BaseKeeper object with a given codec, dedicated
// store key, an AccountKeeper implementation, and a parameter Subspace used to
// store and fetch module parameters. The BaseKeeper also accepts a
// blocklist map. This blocklist describes the set of addresses that are not allowed
// to receive funds through direct and explicit actions, for example, by using a MsgSend or
// by using a SendCoinsFromModuleToAccount execution.
func NewBaseKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ak types.AccountKeeper,
	paramSpace paramtypes.Subspace,
	blockedAddrs map[string]bool,
) BaseKeeper {
	_ = "STUB: not implemented"

	// set KeyTable if it has not already been set
	return *new(BaseKeeper)
}

func NewBaseKeeperWithDeferredCache(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ak types.AccountKeeper,
	paramSpace paramtypes.Subspace,
	blockedAddrs map[string]bool,
	deferredCacheStoreKey sdk.StoreKey,
) BaseKeeper {
	_ = "STUB: not implemented"

	// set KeyTable if it has not already been set
	return *new(BaseKeeper)
}

// WithMintCoinsRestriction restricts the bank Keeper used within a specific module to
// have restricted permissions on minting via function passed in parameter.
// Previous restriction functions can be nested as such:
//
//	bankKeeper.WithMintCoinsRestriction(restriction1).WithMintCoinsRestriction(restriction2)
func (k BaseKeeper) WithMintCoinsRestriction(check MintingRestrictionFn) BaseKeeper {
	_ = "STUB: not implemented"
	return *new(BaseKeeper)
}

// DelegateCoins performs delegation by deducting amt coins from an account with
// address addr. For vesting accounts, delegations amounts are tracked for both
// vesting and vested coins. The coins are then transferred from the delegator
// address to a ModuleAccount address. If any of the delegation amounts are negative,
// an error is returned.
func (k BaseKeeper) DelegateCoins(ctx sdk.Context, delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// emit coin spent event

// UndelegateCoins performs undelegation by crediting amt coins to an account with
// address addr. For vesting accounts, undelegation amounts are tracked for both
// vesting and vested coins. The coins are then transferred from a ModuleAccount
// address to the delegator address. If any of the undelegation amounts are
// negative, an error is returned.
func (k BaseKeeper) UndelegateCoins(ctx sdk.Context, moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSupply retrieves the Supply from store
func (k BaseKeeper) GetSupply(ctx sdk.Context, denom string) sdk.Coin {
	_ = "STUB: not implemented"
	return *new(sdk.Coin)
}

// HasSupply checks if the supply coin exists in store.
func (k BaseKeeper) HasSupply(ctx sdk.Context, denom string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetDenomMetaData retrieves the denomination metadata. returns the metadata and true if the denom exists,
// false otherwise.
func (k BaseKeeper) GetDenomMetaData(ctx sdk.Context, denom string) (types.Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(types.Metadata), false
}

// SetDenomMetaData sets the denominations metadata
func (k BaseKeeper) SetDenomMetaData(ctx sdk.Context, denomMetaData types.Metadata) {
	_ = "STUB: not implemented"
	return
}

// SendCoinsFromModuleToAccount transfers coins from a ModuleAccount to an AccAddress.
// It will panic if the module account does not exist. An error is returned if
// the recipient address is black-listed or if sending the tokens fails.
func (k BaseKeeper) SendCoinsFromModuleToAccount(
	ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SendCoinsFromModuleToModule transfers coins from a ModuleAccount to another.
// It will panic if either module account does not exist.
func (k BaseKeeper) SendCoinsFromModuleToModule(
	ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins,
) error {
	_ = "STUB: not implemented"
	return nil
}

// SendCoinsFromAccountToModule transfers coins from an AccAddress to a ModuleAccount.
// It will panic if the module account does not exist.
func (k BaseKeeper) SendCoinsFromAccountToModule(
	ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins,
) error {
	_ = "STUB: not implemented"
	return nil
}

// DeferredSendCoinsFromAccountToModule transfers coins from an AccAddress to a ModuleAccount.
// It deducts the balance from an accAddress and stores the balance in a mapping for ModuleAccounts.
// In the EndBlocker, it will then perform one deposit for each module account.
// It will panic if the module account does not exist.
func (k BaseKeeper) DeferredSendCoinsFromAccountToModule(
	ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amount sdk.Coins,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Deducts Fees from the Sender Account

// get recipient module address

// get txIndex

//nolint:gosec

// WriteDeferredDepositsToModuleAccounts Iterates on all the deferred deposits and deposit them into the store
func (k BaseKeeper) WriteDeferredBalances(ctx sdk.Context) []abci.Event {
	_ = "STUB: not implemented"
	return nil
}

// maps between bech32 stringified module account address and balance

// slice of modules to be sorted for consistent write order later

// iterate over deferred cache and accumulate totals per module

// add to list of modules

// set the map value

// add to currCoins

// update map

// sort module list

// iterate through module list and add the balance to module bank balances in sorted order

// clear deferred cache

func (k BaseKeeper) IterateDeferredBalances(ctx sdk.Context, cb func(addr sdk.AccAddress, coin sdk.Coin) bool) {
	_ = "STUB: not implemented"
	return
}

// pass cb to deferred cache iterator

// DelegateCoinsFromAccountToModule delegates coins and transfers them from a
// delegator account to a module account. It will panic if the module account
// does not exist or is unauthorized.
func (k BaseKeeper) DelegateCoinsFromAccountToModule(
	ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UndelegateCoinsFromModuleToAccount undelegates the unbonding coins and transfers
// them from a module account to the delegator account. It will panic if the
// module account does not exist or is unauthorized.
func (k BaseKeeper) UndelegateCoinsFromModuleToAccount(
	ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (k BaseKeeper) createCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins, addFn AddFn) error {
	_ = "STUB: not implemented"
	return nil
}

// emit mint event

// MintCoins creates new coins from thin air and adds it to the module account.
// It will panic if the module account does not exist or is unauthorized.
func (k BaseKeeper) MintCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

func (k BaseKeeper) destroyCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins, subFn SubFn) error {
	_ = "STUB: not implemented"
	return nil
}

// emit burn event

// BurnCoins burns coins deletes coins from the balance of the module account.
// It will panic if the module account does not exist or is unauthorized.
func (k BaseKeeper) BurnCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// SetSupply sets the supply for the given coin
func (k BaseKeeper) SetSupply(ctx sdk.Context, coin sdk.Coin) { _ = "STUB: not implemented"; return }

// Bank invariants and IBC requires to remove zero coins.

// trackDelegation tracks the delegation of the given account if it is a vesting account
func (k BaseKeeper) trackDelegation(ctx sdk.Context, addr sdk.AccAddress, balance, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: return error on account.TrackDelegation

// trackUndelegation trakcs undelegation of the given account if it is a vesting account
func (k BaseKeeper) trackUndelegation(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: return error on account.TrackUndelegation

func (k BaseKeeper) GetStoreKey() sdk.StoreKey {
	_ = "STUB: not implemented"

	// for testing
	return *new(sdk.StoreKey)
}

func (k BaseKeeper) GetCdc() codec.BinaryCodec {
	_ = "STUB: not implemented"
	return *new(codec.BinaryCodec)
}
