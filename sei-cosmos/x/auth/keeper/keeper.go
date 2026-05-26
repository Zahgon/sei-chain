package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	paramtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/types"
)

// AccountKeeperI is the interface contract that x/auth's keeper implements.
type AccountKeeperI interface {
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI

	// Return a new account with the next account number. Does not save the new account to the store.
	NewAccount(sdk.Context, types.AccountI) types.AccountI

	// Check if an account exists in the store.
	HasAccount(sdk.Context, sdk.AccAddress) bool

	// Retrieve an account from the store.
	GetAccount(sdk.Context, sdk.AccAddress) types.AccountI

	// Set an account in the store.
	SetAccount(sdk.Context, types.AccountI)

	// Remove an account from the store.
	RemoveAccount(sdk.Context, types.AccountI)

	// Iterate over all accounts, calling the provided function. Stop iteration when it returns true.
	IterateAccounts(sdk.Context, func(types.AccountI) bool)

	// Fetch the public key of an account at a specified address
	GetPubKey(sdk.Context, sdk.AccAddress) (cryptotypes.PubKey, error)

	// Fetch the sequence of an account at a specified address.
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)

	// Fetch the next account number, and increment the internal counter.
	GetNextAccountNumber(sdk.Context) uint64
}

// AccountKeeper encodes/decodes accounts using the go-amino (binary)
// encoding/decoding library.
type AccountKeeper struct {
	key           sdk.StoreKey
	cdc           codec.BinaryCodec
	paramSubspace paramtypes.Subspace
	permAddrs     map[string]types.PermissionsForAddress

	// The prototypical AccountI constructor.
	proto func() types.AccountI
}

var _ AccountKeeperI = &AccountKeeper{}

// NewAccountKeeper returns a new AccountKeeperI that uses go-amino to
// (binary) encode and decode concrete sdk.Accounts.
// `maccPerms` is a map that takes accounts' addresses as keys, and their respective permissions as values. This map is used to construct
// types.PermissionsForAddress and is used in keeper.ValidatePermissions. Permissions are plain strings,
// and don't have to fit into any predefined structure. This auth module does not use account permissions internally, though other modules
// may use auth.Keeper to access the accounts permissions map.
func NewAccountKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramstore paramtypes.Subspace, proto func() types.AccountI,
	maccPerms map[string][]string,
) AccountKeeper {
	_ = "STUB: not implemented"

	// set KeyTable if it has not already been set
	return *new(AccountKeeper)
}

// GetPubKey Returns the PubKey of the account at address
func (ak AccountKeeper) GetPubKey(ctx sdk.Context, addr sdk.AccAddress) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// GetSequence Returns the Sequence of the account at address
func (ak AccountKeeper) GetSequence(ctx sdk.Context, addr sdk.AccAddress) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetNextAccountNumber returns and increments the global account number counter.
// If the global account number is not set, it initializes it with value 0.
func (ak AccountKeeper) GetNextAccountNumber(ctx sdk.Context) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// initialize the account numbers

func (ak AccountKeeper) GetGlobalAccountNumberBytes(ctx sdk.Context) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (ak AccountKeeper) SetGlobalAccountNumberBytes(ctx sdk.Context, bz []byte) {
	_ = "STUB: not implemented"
	return
}

// ValidatePermissions validates that the module account has been granted
// permissions within its set of allowed permissions.
func (ak AccountKeeper) ValidatePermissions(macc types.ModuleAccountI) error {
	_ = "STUB: not implemented"
	return nil
}

// GetModuleAddress returns an address based on the module name
func (ak AccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress)
}

// GetModuleAddressAndPermissions returns an address and permissions based on the module name
func (ak AccountKeeper) GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// GetModuleAccountAndPermissions gets the module account from the auth account store and its
// registered permissions
func (ak AccountKeeper) GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (types.ModuleAccountI, []string) {
	_ = "STUB: not implemented"
	return *new(types.ModuleAccountI), nil
}

// create a new module account

// set the account number

// GetModuleAccount gets the module account from the auth account store, if the account does not
// exist in the AccountKeeper, then it is created.
func (ak AccountKeeper) GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI {
	_ = "STUB: not implemented"
	return *new(types.ModuleAccountI)
}

// SetModuleAccount sets the module account to the auth account store
func (ak AccountKeeper) SetModuleAccount(ctx sdk.Context, macc types.ModuleAccountI) {
	_ = "STUB: not implemented"
	return
}

func (ak AccountKeeper) decodeAccount(bz []byte) types.AccountI {
	_ = "STUB: not implemented"
	return *new(types.AccountI)
}

// MarshalAccount protobuf serializes an Account interface
func (ak AccountKeeper) MarshalAccount(accountI types.AccountI) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalAccount returns an Account interface from raw encoded account
// bytes of a Proto-based Account type
func (ak AccountKeeper) UnmarshalAccount(bz []byte) (types.AccountI, error) {
	_ = "STUB: not implemented"
	return *new(types.AccountI), nil
}

// GetCodec return codec.Codec object used by the keeper
func (ak AccountKeeper) GetCodec() codec.BinaryCodec {
	_ = "STUB: not implemented"
	return *new(codec.BinaryCodec)
}
