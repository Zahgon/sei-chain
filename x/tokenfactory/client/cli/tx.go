package cli

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	banktypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
)

const (
	FlagAllowList            = "allow-list"
	FlagAllowListDescription = "Path to the allow list JSON file with an array of addresses " +
		"that are allowed to send/receive the token. The file should have the following format: {\"addresses\": " +
		"[\"addr1\", \"addr2\"]}, where addr1 and addr2 are bech32 Sei native addresses or EVM addresses."
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCreateDenomCmd broadcast MsgCreateDenom
func NewCreateDenomCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// only parse allow list if it is provided

// Parse the allow list

// NewUpdateDenomCmd broadcast MsgUpdateDenom
func NewUpdateDenomCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Fail if allow list is not provided as this is the only feature that can be updated for now

// NewMintCmd broadcast MsgMint
func NewMintCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewBurnCmd broadcast MsgBurn
func NewBurnCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewChangeAdminCmd broadcast MsgChangeAdmin
func NewChangeAdminCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func NewSetDenomMetadataCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func ParseMetadataJSON(cdc *codec.LegacyAmino, metadataFile string) (banktypes.Metadata, error) {
	_ = "STUB: not implemented"
	return *new(banktypes.Metadata), nil
}

func ParseAllowListJSON(allowListFile string, queryClient evmtypes.QueryClient) (banktypes.AllowList, error) {
	_ = "STUB: not implemented"
	return *new(banktypes.AllowList), nil
}

// Skip duplicate addresses
