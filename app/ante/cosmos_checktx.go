package ante

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/keeper"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/legacy/legacytx"
	authtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/authz"
	bankkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/keeper"
	feegrantkeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant/keeper"
	paramskeeper "github.com/sei-protocol/sei-chain/sei-cosmos/x/params/keeper"
	ibckeeper "github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/keeper"
	tmproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/types"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	oraclekeeper "github.com/sei-protocol/sei-chain/x/oracle/keeper"
)

const maxNestedMsgs = 5

var (
	_ GasTx = (*legacytx.StdTx)(nil) // assert StdTx implements GasTx
)

// GasTx defines a Tx with a GetGas() method which is needed to use SetUpContextDecorator
type GasTx interface {
	sdk.Tx
	GetGas() uint64
}

type HasExtensionOptionsTx interface {
	GetExtensionOptions() []*codectypes.Any
	GetNonCriticalExtensionOptions() []*codectypes.Any
}

// TxWithTimeoutHeight defines the interface a tx must implement in order for
// TxHeightTimeoutDecorator to process the tx.
type TxWithTimeoutHeight interface {
	sdk.Tx

	GetTimeoutHeight() uint64
}

func CosmosCheckTxAnte(
	ctx sdk.Context,
	txConfig client.TxConfig,
	tx sdk.Tx,
	pk paramskeeper.Keeper,
	oraclek oraclekeeper.Keeper,
	ek *evmkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	feegrantKeeper *feegrantkeeper.Keeper,
	ibcKeeper *ibckeeper.Keeper,
) (returnCtx sdk.Context, returnErr error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

func HandleOutofGas(recoveredErr any, gasLimit uint64, gasConsumed uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func CosmosStatelessChecks(tx sdk.Tx, height int64, consensusParams *tmproto.ConsensusParams) (
	isOracleVote bool, err error,
) {
	_ = "STUB: not implemented"
	return false, nil
}

// If there exists a maximum block gas limit, we must ensure that the tx
// does not exceed it.
//nolint:gosec

//nolint:gosec

// PublicKey was omitted from slice since it has already been set in context

// find nested evm messages

func SetGasMeter(ctx sdk.Context, gasLimit uint64, paramsKeeper paramskeeper.Keeper) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

func CheckAndChargeFees(ctx sdk.Context, tx sdk.Tx, accountKeeper authkeeper.AccountKeeper, bankKeeper bankkeeper.Keeper, feegrantKeeper *feegrantkeeper.Keeper, paramsKeeper paramskeeper.Keeper, isGasless bool) (priority int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Determine the required fees by multiplying each required minimum gas
// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
//nolint:gosec // G115: gas is bounded by block gas limit, cannot overflow int64

//nolint:gosec

func chargeFees(ctx sdk.Context, tx sdk.Tx, feeCoins sdk.Coins, accountKeeper authkeeper.AccountKeeper, bankKeeper bankkeeper.Keeper, feegrantKeeper *feegrantkeeper.Keeper) (sdk.AccAddress, error) {
	_ = "STUB: not implemented"
	return *new(sdk.AccAddress), nil
}

// if feegranter set deduct fee from feegranter account.
// this works with only when feegrant enabled.

// deduct the fees

func DecoratePriority(ctx sdk.Context, priority int64, oracleVote bool) sdk.Context {
	_ = "STUB: not implemented"
	return *new(sdk.Context)
}

func CheckMemoLength(tx sdk.Tx, authParams authtypes.Params) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckPubKeys(ctx sdk.Context, tx sdk.Tx, accountKeeper authkeeper.AccountKeeper, authParams authtypes.Params) ([]authtypes.AccountI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

func CheckSignatures(ctx sdk.Context, txConfig client.TxConfig, tx sdk.Tx, signerAccounts []authtypes.AccountI, authParams authtypes.Params) (sdk.Events, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Events), nil
}

// stdSigs contains the sequence number, account number, and signatures.
// When simulating, this would just be a 0-length slice.

// check that signer length and signature length are the same

// CheckTx and ReCheckTx discard these events (see CosmosCheckTxAnte); building them
// still runs SignatureDataToBz + base64 per signer — measurable CPU/alloc on hot path.

// make a SignatureV2 with PubKey filled in from above

// Check account sequence number.

// retrieve signer data

// If all signers are using SIGN_MODE_LEGACY_AMINO, we rely on VerifySignature to check account sequence number,
// and therefore communicate sequence number as a potential cause of error.

func UpdateSigners(ctx sdk.Context, tx sdk.Tx, accountKeeper authkeeper.AccountKeeper, evmKeeper *evmkeeper.Keeper) (sdk.Events, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Events), nil
}

// check if there is non-zero balance

func CheckMessage(ctx sdk.Context, tx sdk.Tx, ibcKeeper *ibckeeper.Keeper, oracleKeeper oraclekeeper.Keeper) error {
	_ = "STUB: not implemented"
	// keep track of total packet messages and number of redundancies across `RecvPacket`, `AcknowledgePacket`, and `TimeoutPacket/OnClose`
	return nil
}

// only return error if all packet messages are redundant

func CheckAuthzContainsEvm(authzMsg *authz.MsgExec, nestedLvl int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check if message type is authz exec or evm

// find nested to check for evm
