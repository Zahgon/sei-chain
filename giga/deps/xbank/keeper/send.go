package keeper

import (
	"github.com/sei-protocol/sei-chain/giga/deps/xbank/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	cosmosbanktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

const (
	TokenFactoryPrefix = "factory"
)

var CoinbaseAddressPrefix = []byte("evm_coinbase")

// SendKeeper defines a module interface that facilitates the transfer of coins
// between accounts without the possibility of creating coins.
type SendKeeper interface {
	ViewKeeper

	InputOutputCoins(ctx sdk.Context, inputs []types.Input, outputs []types.Output) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsWithoutAccCreation(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsAndWei(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amt sdk.Int, wei sdk.Int) error
	SubUnlockedCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins, checkNeg bool) error
	AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins, checkNeg bool) error
	SubWei(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Int) error
	AddWei(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Int) error

	GetParams(ctx sdk.Context) cosmosbanktypes.Params
	SetParams(ctx sdk.Context, params cosmosbanktypes.Params)

	IsSendEnabledCoin(ctx sdk.Context, coin sdk.Coin) bool
	IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error
	SetDenomAllowList(ctx sdk.Context, denom string, allowList types.AllowList)
	GetDenomAllowList(ctx sdk.Context, denom string) types.AllowList
	IsInDenomAllowList(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins, cache map[string]AllowedAddresses) bool

	BlockedAddr(addr sdk.AccAddress) bool
	RegisterRecipientChecker(RecipientChecker)
}

type RecipientChecker = func(ctx sdk.Context, recipient sdk.AccAddress) bool

var _ SendKeeper = (*BaseSendKeeper)(nil)
var OneUseiInWei sdk.Int = sdk.NewInt(1_000_000_000_000)

// BaseSendKeeper only allows transfers between accounts without the possibility of
// creating coins. It implements the SendKeeper interface.
type BaseSendKeeper struct {
	BaseViewKeeper

	cdc        codec.BinaryCodec
	ak         types.AccountKeeper
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace

	// list of addresses that are restricted from receiving transactions
	blockedAddrs      map[string]bool
	recipientCheckers *[]RecipientChecker
}

func NewBaseSendKeeper(
	cdc codec.BinaryCodec, storeKey sdk.StoreKey, ak types.AccountKeeper, paramSpace paramtypes.Subspace, blockedAddrs map[string]bool,
) BaseSendKeeper {
	_ = "STUB: not implemented"
	return *new(BaseSendKeeper)
}

// GetParams returns the total set of bank parameters.
func (k BaseSendKeeper) GetParams(ctx sdk.Context) (params cosmosbanktypes.Params) {
	_ = "STUB: not implemented"
	return *new(cosmosbanktypes.Params)
}

// SetParams sets the total set of bank parameters.
func (k BaseSendKeeper) SetParams(ctx sdk.Context, params cosmosbanktypes.Params) {
	_ = "STUB: not implemented"
	return
}

// InputOutputCoins performs multi-send functionality. It accepts a series of
// inputs that correspond to a series of outputs. It returns an error if the
// inputs and outputs don't lineup or if any single transfer of tokens fails.
func (k BaseSendKeeper) InputOutputCoins(ctx sdk.Context, inputs []types.Input, outputs []types.Output) error {
	_ = "STUB: not implemented"
	// Safety check ensuring that when sending coins the keeper must maintain the
	// Check supply invariant and validity of Coins.
	return nil
}

// Create account if recipient does not exist.
//
// NOTE: This should ultimately be removed in favor a more flexible approach
// such as delegated fee messages.

// SendCoins transfers amt coins from a sending account to a receiving account.
// An error is returned upon failure.
func (k BaseSendKeeper) SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// Create account if recipient does not exist.
//
// NOTE: This should ultimately be removed in favor a more flexible approach
// such as delegated fee messages.

func (k BaseSendKeeper) SendCoinsWithoutAccCreation(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

func (k BaseSendKeeper) sendCoinsWithoutAccCreation(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins, checkNeg bool) error {
	_ = "STUB: not implemented"
	return nil
}

// SubUnlockedCoins removes the unlocked amt coins of the given account. An error is
// returned if the resulting balance is negative or the initial amount is invalid.
// A coin_spent event is emitted after.
func (k BaseSendKeeper) SubUnlockedCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins, checkNeg bool) error {
	_ = "STUB: not implemented"
	return nil
}

// emit coin spent event

// AddCoins increase the addr balance by the given amt. Fails if the provided amt is invalid.
// It emits a coin received event.
func (k BaseSendKeeper) AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins, checkNeg bool) error {
	_ = "STUB: not implemented"
	return nil
}

// emit coin received event

// initBalances sets the balance (multiple coins) for an account by address.
// An error is returned upon failure.
func (k BaseSendKeeper) initBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) error {
	_ = "STUB: not implemented"
	return nil
}

// Bank invariants require to not store zero balances.

// setBalance sets the coin balance for an account by address.
func (k BaseSendKeeper) setBalance(ctx sdk.Context, addr sdk.AccAddress, balance sdk.Coin, checkNeg bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Bank invariants require to not store zero balances.

func (k BaseSendKeeper) setWeiBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// IsSendEnabledCoins checks the coins provide and returns an ErrSendDisabled if
// any of the coins are not configured for sending.  Returns nil if sending is enabled
// for all provided coin
func (k BaseSendKeeper) IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error {
	_ = "STUB: not implemented"
	return nil
}

// IsSendEnabledCoin returns the current SendEnabled status of the provided coin's denom
func (k BaseSendKeeper) IsSendEnabledCoin(ctx sdk.Context, coin sdk.Coin) bool {
	_ = "STUB: not implemented"
	return false
}

// BlockedAddr checks if a given address is restricted from
// receiving funds.
func (k BaseSendKeeper) BlockedAddr(addr sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

func (k BaseSendKeeper) SubWei(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// no need to change usei balance

func (k BaseSendKeeper) AddWei(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// no need to change usei balance

func (k BaseSendKeeper) SendCoinsAndWei(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amt sdk.Int, wei sdk.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func (k BaseSendKeeper) RegisterRecipientChecker(rc RecipientChecker) {
	_ = "STUB: not implemented"
	return
}

func (k BaseSendKeeper) CanSendTo(ctx sdk.Context, recipient sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

func SplitUseiWeiAmount(amt sdk.Int) (sdk.Int, sdk.Int) {
	_ = "STUB: not implemented"
	return *new(sdk.Int), *new(sdk.Int)
}

func (k BaseSendKeeper) SetDenomAllowList(ctx sdk.Context, denom string, allowList types.AllowList) {
	_ = "STUB: not implemented"
	return
}

func (k BaseSendKeeper) GetDenomAllowList(ctx sdk.Context, denom string) types.AllowList {
	_ = "STUB: not implemented"
	return *new(types.AllowList)
}

// IsInDenomAllowList checks if the given address is allowed to send the given coins.
// The check is performed only fot token factory denoms. For each token factory denom,
// it checks if there is allow list for the given denom. If there is no allow list,
// the address is allowed to send the coins. If there is an allow list, the address is
// allowed to send the coins only if it is in the allow list.
func (k BaseSendKeeper) IsInDenomAllowList(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins, cache map[string]AllowedAddresses) bool {
	_ = "STUB: not implemented"
	return false

	// Skip if denom does not contain the token factory prefix
}

// skip if there is no allow list for the denom

func (k BaseSendKeeper) getAllowedAddresses(ctx sdk.Context, cache map[string]AllowedAddresses, denom string) AllowedAddresses {
	_ = "STUB: not implemented"
	return *new(AllowedAddresses)
}

// we cache even if the allowList is empty to avoid multiple db reads

type AllowedAddresses struct {
	set map[string]struct{}
}

func (a AllowedAddresses) contains(address sdk.AccAddress) bool {
	_ = "STUB: not implemented"
	return false
}

func (k BaseSendKeeper) buildAllowedAddressesMap(allowList types.AllowList) AllowedAddresses {
	_ = "STUB: not implemented"
	return *new(AllowedAddresses)
}
