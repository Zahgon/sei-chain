package baseapp

import (
	"time"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// Paramspace defines the parameter subspace to be used for the paramstore.
const Paramspace = "baseapp"

// Parameter store keys for all the consensus parameter types.
var (
	ParamStoreKeyEvidenceParams  = []byte("EvidenceParams")
	ParamStoreKeyValidatorParams = []byte("ValidatorParams")
	ParamStoreKeyBlockParams     = []byte("BlockParams")
	ParamStoreKeyVersionParams   = []byte("VersionParams")
	ParamStoreKeySynchronyParams = []byte("SynchronyParams")
	ParamStoreKeyTimeoutParams   = []byte("TimeoutParams")
	ParamStoreKeyABCIParams      = []byte("ABCIParams")
)

// ParamStore defines the interface the parameter store used by the BaseApp must
// fulfill.
type ParamStore interface {
	Get(ctx sdk.Context, key []byte, ptr interface{})
	Has(ctx sdk.Context, key []byte) bool
	Set(ctx sdk.Context, key []byte, param interface{})
}

// ValidateBlockParams defines a stateless validation on BlockParams. This function
// is called whenever the parameters are updated or stored.
func ValidateBlockParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

// ValidateEvidenceParams defines a stateless validation on EvidenceParams. This
// function is called whenever the parameters are updated or stored.
func ValidateEvidenceParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

// ValidateValidatorParams defines a stateless validation on ValidatorParams. This
// function is called whenever the parameters are updated or stored.
func ValidateValidatorParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

func ValidateVersionParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

func ValidateSynchronyParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

func ValidateTimeoutParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

func ValidateABCIParams(i interface{}) error { _ = "STUB: not implemented"; return nil }

func validateDurationPointer(i *time.Duration, name string) error {
	_ = "STUB: not implemented"
	return nil
}
