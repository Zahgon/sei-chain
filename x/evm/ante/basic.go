package ante

import (
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

type BasicDecorator struct {
	k *keeper.Keeper
}

func NewBasicDecorator(k *keeper.Keeper) *BasicDecorator { _ = "STUB: not implemented"; return nil }

// cherrypicked from go-ethereum:txpool:ValidateTransaction
func (gl BasicDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_ = "STUB: not implemented"
	return *new(sdk.Context), nil
}

// bump nonce if it is for some reason not incremented (e.g. ante failure)

// Check if gas exceed the limit

// If there exists a maximum block gas limit, we must ensure that the tx
// does not exceed it.
//nolint:gosec

//TODO: support blobs (leaving this commented out)
// Ensure blob transactions have valid commitments
//if etx.Type() == ethtypes.BlobTxType {
//	sidecar := etx.BlobTxSidecar()
//	if sidecar == nil {
//		return ctx, fmt.Errorf("missing sidecar in blob transaction")
//	}
//	// Ensure the number of items in the blob transaction and various side
//	// data match up before doing any expensive validations
//	hashes := etx.BlobHashes()
//	if len(hashes) == 0 {
//		return ctx, fmt.Errorf("blobless blob transaction")
//	}
//	if len(hashes) > params.MaxBlobGasPerBlock/params.BlobTxBlobGasPerBlob {
//		return ctx, fmt.Errorf("too many blobs in transaction: have %d, permitted %d", len(hashes), params.MaxBlobGasPerBlock/params.BlobTxBlobGasPerBlob)
//	}
//	if err := validateBlobSidecar(hashes, sidecar); err != nil {
//		return ctx, err
//	}
//}

// NOTE: if we re-enable blob support we can put this back.
// this allows code coverage to calculate correctly
//func validateBlobSidecar(hashes []common.Hash, sidecar *ethtypes.BlobTxSidecar) error {
//	if len(sidecar.Blobs) != len(hashes) {
//		return fmt.Errorf("invalid number of %d blobs compared to %d blob hashes", len(sidecar.Blobs), len(hashes))
//	}
//	if len(sidecar.Commitments) != len(hashes) {
//		return fmt.Errorf("invalid number of %d blob commitments compared to %d blob hashes", len(sidecar.Commitments), len(hashes))
//	}
//	if len(sidecar.Proofs) != len(hashes) {
//		return fmt.Errorf("invalid number of %d blob proofs compared to %d blob hashes", len(sidecar.Proofs), len(hashes))
//	}
//	// Blob quantities match up, validate that the provers match with the
//	// transaction hash before getting to the cryptography
//	hasher := sha256.New()
//	for i, want := range hashes {
//		hasher.Write(sidecar.Commitments[i][:])
//		hash := hasher.Sum(nil)
//		hasher.Reset()
//
//		var vhash common.Hash
//		vhash[0] = params.BlobTxHashVersion
//		copy(vhash[1:], hash[1:])
//
//		if vhash != want {
//			return fmt.Errorf("blob %d: computed hash %#x mismatches transaction one %#x", i, vhash, want)
//		}
//	}
//	// Blob commitments match with the hashes in the transaction, verify the
//	// blobs themselves via KZG
//	for i := range sidecar.Blobs {
//		if err := kzg4844.VerifyBlobProof(sidecar.Blobs[i], sidecar.Commitments[i], sidecar.Proofs[i]); err != nil {
//			return fmt.Errorf("invalid blob %d: %v", i, err)
//		}
//	}
//	return nil
//}
