package keeper

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/exported"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/evidence/types"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

// Keeper defines the evidence module's keeper. The keeper is responsible for
// managing persistence, state transitions and query handling for the evidence
// module.
type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       sdk.StoreKey
	router         types.Router
	stakingKeeper  types.StakingKeeper
	slashingKeeper types.SlashingKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, storeKey sdk.StoreKey, stakingKeeper types.StakingKeeper,
	slashingKeeper types.SlashingKeeper,
) *Keeper {
	_ = "STUB: not implemented"
	return nil
}

// SetRouter sets the Evidence Handler router for the x/evidence module. Note,
// we allow the ability to set the router after the Keeper is constructed as a
// given Handler may need access the Keeper before being constructed. The router
// may only be set once and will be sealed if it's not already sealed.
func (k *Keeper) SetRouter(rtr types.Router) {
	_ = "STUB: not implemented"
	// It is vital to seal the Evidence Handler router as to not allow further
	// handlers to be registered after the keeper is created since this
	// could create invalid or non-deterministic behavior.
	return
}

// GetEvidenceHandler returns a registered Handler for a given Evidence type. If
// no handler exists, an error is returned.
func (k Keeper) GetEvidenceHandler(evidenceRoute string) (types.Handler, error) {
	_ = "STUB: not implemented"
	return *new(types.Handler), nil
}

// SubmitEvidence attempts to match evidence against the keepers router and execute
// the corresponding registered Evidence Handler. An error is returned if no
// registered Handler exists or if the Handler fails. Otherwise, the evidence is
// persisted.
func (k Keeper) SubmitEvidence(ctx sdk.Context, evidence exported.Evidence) error {
	_ = "STUB: not implemented"
	return nil
}

// SetEvidence sets Evidence by hash in the module's KVStore.
func (k Keeper) SetEvidence(ctx sdk.Context, evidence exported.Evidence) {
	_ = "STUB: not implemented"
	return
}

// GetEvidence retrieves Evidence by hash if it exists. If no Evidence exists for
// the given hash, (nil, false) is returned.
func (k Keeper) GetEvidence(ctx sdk.Context, hash tmbytes.HexBytes) (exported.Evidence, bool) {
	_ = "STUB: not implemented"
	return *new(exported.Evidence), false
}

// IterateEvidence provides an interator over all stored Evidence objects. For
// each Evidence object, cb will be called. If the cb returns true, the iterator
// will close and stop.
func (k Keeper) IterateEvidence(ctx sdk.Context, cb func(exported.Evidence) bool) {
	_ = "STUB: not implemented"
	return
}

// GetAllEvidence returns all stored Evidence objects.
func (k Keeper) GetAllEvidence(ctx sdk.Context) (evidence []exported.Evidence) {
	_ = "STUB: not implemented"
	return nil
}

// MustUnmarshalEvidence attempts to decode and return an Evidence object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalEvidence(bz []byte) exported.Evidence {
	_ = "STUB: not implemented"
	return *new(exported.Evidence)
}

// MustMarshalEvidence attempts to encode an Evidence object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalEvidence(evidence exported.Evidence) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MarshalEvidence protobuf serializes an Evidence interface
func (k Keeper) MarshalEvidence(evidenceI exported.Evidence) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalEvidence returns an Evidence interface from raw encoded evidence
// bytes of a Proto-based Evidence type
func (k Keeper) UnmarshalEvidence(bz []byte) (exported.Evidence, error) {
	_ = "STUB: not implemented"
	return *new(exported.Evidence), nil
}
