package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

var _ types.ContractOpsKeeper = PermissionedKeeper{}

// decoratedKeeper contains a subset of the wasm keeper that are already or can be guarded by an authorization policy in the future
type decoratedKeeper interface {
	create(ctx sdk.Context, creator sdk.AccAddress, wasmCode []byte, instantiateAccess *types.AccessConfig, authZ AuthorizationPolicy) (codeID uint64, err error)
	instantiate(ctx sdk.Context, codeID uint64, creator, admin sdk.AccAddress, initMsg []byte, label string, deposit sdk.Coins, authZ AuthorizationPolicy) (sdk.AccAddress, []byte, error)
	migrate(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, newCodeID uint64, msg []byte, authZ AuthorizationPolicy) ([]byte, error)
	setContractAdmin(ctx sdk.Context, contractAddress, caller, newAdmin sdk.AccAddress, authZ AuthorizationPolicy) error
	pinCode(ctx sdk.Context, codeID uint64) error
	unpinCode(ctx sdk.Context, codeID uint64) error
	execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
	setContractInfoExtension(ctx sdk.Context, contract sdk.AccAddress, extra types.ContractInfoExtension) error
	setAccessConfig(ctx sdk.Context, codeID uint64, config types.AccessConfig) error
}

type PermissionedKeeper struct {
	authZPolicy AuthorizationPolicy
	nested      decoratedKeeper
}

func NewPermissionedKeeper(nested decoratedKeeper, authZPolicy AuthorizationPolicy) *PermissionedKeeper {
	_ = "STUB: not implemented"
	return nil
}

func NewGovPermissionKeeper(nested decoratedKeeper) *PermissionedKeeper {
	_ = "STUB: not implemented"
	return nil
}

func NewDefaultPermissionKeeper(nested decoratedKeeper) *PermissionedKeeper {
	_ = "STUB: not implemented"
	return nil
}

func (p PermissionedKeeper) Create(ctx sdk.Context, creator sdk.AccAddress, wasmCode []byte, instantiateAccess *types.AccessConfig) (codeID uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p PermissionedKeeper) Instantiate(ctx sdk.Context, codeID uint64, creator, admin sdk.AccAddress, initMsg []byte, label string, deposit sdk.Coins) (sdk.AccAddress, []byte, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil, nil
}

func (p PermissionedKeeper) Execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PermissionedKeeper) Migrate(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, newCodeID uint64, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PermissionedKeeper) Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PermissionedKeeper) UpdateContractAdmin(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, newAdmin sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PermissionedKeeper) ClearContractAdmin(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PermissionedKeeper) PinCode(ctx sdk.Context, codeID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PermissionedKeeper) UnpinCode(ctx sdk.Context, codeID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// SetExtraContractAttributes updates the extra attributes that can be stored with the contract info
func (p PermissionedKeeper) SetContractInfoExtension(ctx sdk.Context, contract sdk.AccAddress, extra types.ContractInfoExtension) error {
	_ = "STUB: not implemented"
	return nil
}

// SetAccessConfig updates the access config of a code id.
func (p PermissionedKeeper) SetAccessConfig(ctx sdk.Context, codeID uint64, config types.AccessConfig) error {
	_ = "STUB: not implemented"
	return nil
}
