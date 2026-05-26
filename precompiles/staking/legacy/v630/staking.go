package v630

import (
	"embed"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common/legacy/v630"
	"github.com/sei-protocol/sei-chain/precompiles/utils"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	stakingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/staking/types"
	"github.com/sei-protocol/seilog"
)

var logger = seilog.NewLogger("precompiles", "staking", "legacy", "v630")

const (
	DelegateMethod                      = "delegate"
	RedelegateMethod                    = "redelegate"
	UndelegateMethod                    = "undelegate"
	DelegationMethod                    = "delegation"
	CreateValidatorMethod               = "createValidator"
	EditValidatorMethod                 = "editValidator"
	ValidatorsMethod                    = "validators"
	ValidatorMethod                     = "validator"
	ValidatorDelegationsMethod          = "validatorDelegations"
	ValidatorUnbondingDelegationsMethod = "validatorUnbondingDelegations"
	UnbondingDelegationMethod           = "unbondingDelegation"
	DelegatorDelegationsMethod          = "delegatorDelegations"
	DelegatorValidatorMethod            = "delegatorValidator"
	DelegatorUnbondingDelegationsMethod = "delegatorUnbondingDelegations"
	RedelegationsMethod                 = "redelegations"
	DelegatorValidatorsMethod           = "delegatorValidators"
	HistoricalInfoMethod                = "historicalInfo"
	PoolMethod                          = "pool"
	ParamsMethod                        = "params"
)

const (
	StakingAddress = "0x0000000000000000000000000000000000001005"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

type PrecompileExecutor struct {
	stakingKeeper  utils.StakingKeeper
	stakingQuerier utils.StakingQuerier
	evmKeeper      utils.EVMKeeper
	bankKeeper     utils.BankKeeper
	address        common.Address

	DelegateID                      []byte
	RedelegateID                    []byte
	UndelegateID                    []byte
	DelegationID                    []byte
	CreateValidatorID               []byte
	EditValidatorID                 []byte
	ValidatorsID                    []byte
	ValidatorID                     []byte
	ValidatorDelegationsID          []byte
	ValidatorUnbondingDelegationsID []byte
	UnbondingDelegationID           []byte
	DelegatorDelegationsID          []byte
	DelegatorValidatorID            []byte
	DelegatorUnbondingDelegationsID []byte
	RedelegationsID                 []byte
	DelegatorValidatorsID           []byte
	HistoricalInfoID                []byte
	PoolID                          []byte
	ParamsID                        []byte
}

func NewPrecompile(keepers utils.Keepers) (*pcommon.DynamicGasPrecompile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p PrecompileExecutor) EVMKeeper() utils.EVMKeeper {
	_ = "STUB: not implemented"
	return *new(utils.EVMKeeper)
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM, suppliedGas uint64, hooks *tracing.Hooks) (ret []byte, remainingGas uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) delegate(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// if delegator is associated, then it must have Account set already
// if delegator is not associated, then it can't delegate anyway (since
// there is no good way to merge delegations if it becomes associated)

// Emit EVM event

// Log error but don't fail the transaction

func (p PrecompileExecutor) redelegate(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, evm *vm.EVM) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Emit EVM event

// Log error but don't fail the transaction

func (p PrecompileExecutor) undelegate(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, evm *vm.EVM) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Emit EVM event

// Log error but don't fail the transaction

type Delegation struct {
	Balance    Balance
	Delegation DelegationDetails
}

type Balance struct {
	Amount *big.Int
	Denom  string
}

type DelegationDetails struct {
	DelegatorAddress string
	Shares           *big.Int
	Decimals         *big.Int
	ValidatorAddress string
}

func (p PrecompileExecutor) delegation(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

type ValidatorsResponse struct {
	Validators []Validator
	NextKey    []byte
}

type DelegationsResponse struct {
	Delegations []Delegation
	NextKey     []byte
}

type UnbondingDelegationsResponse struct {
	UnbondingDelegations []UnbondingDelegation
	NextKey              []byte
}

type RedelegationsResponse struct {
	Redelegations []Redelegation
	NextKey       []byte
}

type Validator struct {
	OperatorAddress         string
	ConsensusPubkey         []byte
	Jailed                  bool
	Status                  int32
	Tokens                  string
	DelegatorShares         string
	Description             string
	UnbondingHeight         int64
	UnbondingTime           int64
	CommissionRate          string
	CommissionMaxRate       string
	CommissionMaxChangeRate string
	CommissionUpdateTime    int64
	MinSelfDelegation       string
}

func (p PrecompileExecutor) validators(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) createValidator(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Extract arguments

// Get validator address (caller's associated Sei address)

// Parse public key from hex

// Create ed25519 public key

// Parse commission rates

// Validate minimum self delegation

// Emit EVM event

// Log error but don't fail the transaction

func (p PrecompileExecutor) editValidator(ctx sdk.Context, method *abi.Method, caller common.Address, args []interface{}, value *big.Int, hooks *tracing.Hooks, evm *vm.EVM) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Extract arguments

// Get validator address (caller's associated Sei address)

// Parse commission rate if provided

// Convert min self delegation if not zero

// Emit EVM event

// Log error but don't fail the transaction

func (p PrecompileExecutor) validator(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) validatorDelegations(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) validatorUnbondingDelegations(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) unbondingDelegation(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) delegatorDelegations(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) delegatorValidator(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) delegatorUnbondingDelegations(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) redelegations(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) delegatorValidators(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) historicalInfo(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Use the requested height

func (p PrecompileExecutor) pool(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p PrecompileExecutor) params(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Helper function to convert stakingtypes.Validator to precompile Validator type
func convertValidatorToPrecompileType(val stakingtypes.Validator) Validator {
	_ = "STUB: not implemented"
	return *new(Validator)
}

// Additional types for new query methods
type UnbondingDelegationEntry struct {
	CreationHeight int64
	CompletionTime int64
	InitialBalance string
	Balance        string
}

type UnbondingDelegation struct {
	DelegatorAddress string
	ValidatorAddress string
	Entries          []UnbondingDelegationEntry
}

type RedelegationEntry struct {
	CreationHeight int64
	CompletionTime int64
	InitialBalance string
	SharesDst      string
}

type Redelegation struct {
	DelegatorAddress    string
	ValidatorSrcAddress string
	ValidatorDstAddress string
	Entries             []RedelegationEntry
}

type HistoricalInfo struct {
	Height     int64
	Validators []Validator
}

type Pool struct {
	NotBondedTokens string
	BondedTokens    string
}

type Params struct {
	UnbondingTime                      uint64
	MaxValidators                      uint32
	MaxEntries                         uint32
	HistoricalEntries                  uint32
	BondDenom                          string
	MinCommissionRate                  string
	MaxVotingPowerRatio                string
	MaxVotingPowerEnforcementThreshold string
}
