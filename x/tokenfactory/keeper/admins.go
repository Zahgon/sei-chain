package keeper

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/tokenfactory/types"
)

// GetAuthorityMetadata returns the authority metadata for a specific denom
func (k Keeper) GetAuthorityMetadata(ctx sdk.Context, denom string) (types.DenomAuthorityMetadata, error) {
	_ = "STUB: not implemented"
	return *new(types.DenomAuthorityMetadata), nil
}

// setAuthorityMetadata stores authority metadata for a specific denom
func (k Keeper) setAuthorityMetadata(ctx sdk.Context, denom string, metadata types.DenomAuthorityMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (k Keeper) setAdmin(ctx sdk.Context, denom string, admin string) error {
	_ = "STUB: not implemented"
	return nil
}
