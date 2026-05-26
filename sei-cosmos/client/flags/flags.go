package flags

import (
	tmcli "github.com/sei-protocol/sei-chain/sei-tendermint/libs/cli"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/crypto/keyring"
)

const (
	// DefaultGasAdjustment is applied to gas estimates to avoid tx execution
	// failures due to state changes that might occur between the tx simulation
	// and the actual run.
	DefaultGasAdjustment = 1.0
	DefaultGasLimit      = 200000
	GasFlagAuto          = "auto"

	// DefaultKeyringBackend
	DefaultKeyringBackend = keyring.BackendOS

	// BroadcastBlock defines a tx broadcasting mode where the client waits for
	// the tx to be committed in a block.
	BroadcastBlock = "block"
	// BroadcastSync defines a tx broadcasting mode where the client waits for
	// a CheckTx execution response only.
	BroadcastSync = "sync"
	// BroadcastAsync defines a tx broadcasting mode where the client returns
	// immediately.
	BroadcastAsync = "async"

	// SignModeDirect is the value of the --sign-mode flag for SIGN_MODE_DIRECT
	SignModeDirect = "direct"
	// SignModeLegacyAminoJSON is the value of the --sign-mode flag for SIGN_MODE_LEGACY_AMINO_JSON
	SignModeLegacyAminoJSON = "amino-json"
	// SignModeEIP191 is the value of the --sign-mode flag for SIGN_MODE_EIP_191
	SignModeEIP191 = "eip-191"
)

// List of CLI flags
const (
	FlagHome             = tmcli.HomeFlag
	FlagKeyringDir       = "keyring-dir"
	FlagUseLedger        = "ledger"
	FlagChainID          = "chain-id"
	FlagNode             = "node"
	FlagHeight           = "height"
	FlagGasAdjustment    = "gas-adjustment"
	FlagFrom             = "from"
	FlagName             = "name"
	FlagAccountNumber    = "account-number"
	FlagSequence         = "sequence"
	FlagNote             = "note"
	FlagFees             = "fees"
	FlagGas              = "gas"
	FlagGasPrices        = "gas-prices"
	FlagBroadcastMode    = "broadcast-mode"
	FlagDryRun           = "dry-run"
	FlagGenerateOnly     = "generate-only"
	FlagOffline          = "offline"
	FlagOutputDocument   = "output-document" // inspired by wget -O
	FlagSkipConfirmation = "yes"
	FlagProve            = "prove"
	FlagKeyringBackend   = "keyring-backend"
	FlagPage             = "page"
	FlagLimit            = "limit"
	FlagSignMode         = "sign-mode"
	FlagPageKey          = "page-key"
	FlagOffset           = "offset"
	FlagCountTotal       = "count-total"
	FlagTimeoutHeight    = "timeout-height"
	FlagKeyAlgorithm     = "algo"
	FlagFeeAccount       = "fee-account"
	FlagReverse          = "reverse"
	FlagOutput           = tmcli.OutputFlag

	// Tendermint logging flags
	FlagLogLevel  = "log_level"
	FlagLogFormat = "log_format"
)

// LineBreak can be included in a command list to provide a blank line
// to help with readability
var LineBreak = &cobra.Command{Run: func(*cobra.Command, []string) {}}

// AddQueryFlagsToCmd adds common flags to a module query command.
func AddQueryFlagsToCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// AddTxFlagsToCmd adds common flags to a module tx command.
func AddTxFlagsToCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// --gas can accept integers and "auto"

// AddPaginationFlagsToCmd adds common pagination flags to cmd
func AddPaginationFlagsToCmd(cmd *cobra.Command, query string) { _ = "STUB: not implemented"; return }

// GasSetting encapsulates the possible values passed through the --gas flag.
type GasSetting struct {
	Simulate bool
	Gas      uint64
}

func (v *GasSetting) String() string { _ = "STUB: not implemented"; return "" }

// ParseGasSetting parses a string gas value. The value may either be 'auto',
// which indicates a transaction should be executed in simulate mode to
// automatically find a sufficient gas value, or a string integer. It returns an
// error if a string integer is provided which cannot be parsed.
func ParseGasSetting(gasStr string) (GasSetting, error) {
	_ = "STUB: not implemented"
	return *new(GasSetting), nil
}
