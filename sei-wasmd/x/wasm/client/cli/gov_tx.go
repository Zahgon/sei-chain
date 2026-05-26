package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-wasmd/x/wasm/types"
)

func ProposalStoreCodeCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func ProposalInstantiateContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func ProposalMigrateContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func ProposalExecuteContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func ProposalSudoContractCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flagsExecute

// type values must match the "ProposalHandler" "routes" in cli

func ProposalUpdateContractAdminCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func ProposalClearContractAdminCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func ProposalPinCodesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func parsePinCodesArgs(args []string) ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

func ProposalUnpinCodesCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli

func parseAccessConfig(config string) (types.AccessConfig, error) {
	_ = "STUB: not implemented"
	return *new(types.AccessConfig), nil
}

func parseAccessConfigUpdates(args []string) ([]types.AccessConfigUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// format: code_id,access_config
// access_config: nobody|everybody|address

func ProposalUpdateInstantiateConfigCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// proposal flags

// type values must match the "ProposalHandler" "routes" in cli
