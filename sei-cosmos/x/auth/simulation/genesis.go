package simulation

import (
	"math/rand"

	"github.com/sei-protocol/sei-chain/sei-cosmos/types/module"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/types"
)

// Simulation parameter constants
const (
	MaxMemoChars           = "max_memo_characters"
	TxSigLimit             = "tx_sig_limit"
	TxSizeCostPerByte      = "tx_size_cost_per_byte"
	SigVerifyCostED25519   = "sig_verify_cost_ed25519"
	SigVerifyCostSECP256K1 = "sig_verify_cost_secp256k1"
)

// RandomGenesisAccounts defines the default RandomGenesisAccountsFn used on the SDK.
// It creates a slice of BaseAccount, ContinuousVestingAccount and DelayedVestingAccount.
func RandomGenesisAccounts(simState *module.SimulationState) types.GenesisAccounts {
	_ = "STUB: not implemented"
	return *new(types.GenesisAccounts)
}

// Only consider making a vesting account once the initial bonded validator
// set is exhausted due to needing to track DelegatedVesting.

// Allow for some vesting accounts to vest very quickly while others very slowly.

// GenMaxMemoChars randomized MaxMemoChars
func GenMaxMemoChars(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // always returns a positive value

// GenTxSigLimit randomized TxSigLimit
// make sure that sigLimit is always high
// so that arbitrarily simulated messages from other
// modules can still create valid transactions
func GenTxSigLimit(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // always returns a positive value

// GenTxSizeCostPerByte randomized TxSizeCostPerByte
func GenTxSizeCostPerByte(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // always returns a positive value

// GenSigVerifyCostED25519 randomized SigVerifyCostED25519
func GenSigVerifyCostED25519(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // always returns a positive value

// GenSigVerifyCostSECP256K1 randomized SigVerifyCostSECP256K1
func GenSigVerifyCostSECP256K1(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // always returns a positive value

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simState *module.SimulationState, randGenAccountsFn types.RandomGenesisAccountsFn) {
	_ = "STUB: not implemented"
	return
}
