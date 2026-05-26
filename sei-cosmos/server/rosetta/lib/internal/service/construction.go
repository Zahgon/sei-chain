package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// ConstructionCombine Combine creates a network-specific transaction from an unsigned transaction
// and an array of provided signatures. The signed transaction returned from this method will be
// sent to the /construction/submit endpoint by the caller.
func (on OnlineNetwork) ConstructionCombine(ctx context.Context, request *types.ConstructionCombineRequest) (*types.ConstructionCombineResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionDerive Derive returns the AccountIdentifier associated with a public key.
func (on OnlineNetwork) ConstructionDerive(_ context.Context, request *types.ConstructionDeriveRequest) (*types.ConstructionDeriveResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionHash TransactionHash returns the network-specific transaction hash for a signed
// transaction.
func (on OnlineNetwork) ConstructionHash(ctx context.Context, request *types.ConstructionHashRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionMetadata Get any information required to construct a transaction for a specific
// network (i.e. ChainID, Gas, Memo, ...).
func (on OnlineNetwork) ConstructionMetadata(ctx context.Context, request *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionParse Parse is called on both unsigned and signed transactions to understand the
// intent of the formulated transaction. This is run as a sanity check before signing (after
// /construction/payloads) and before broadcast (after /construction/combine).
func (on OnlineNetwork) ConstructionParse(ctx context.Context, request *types.ConstructionParseRequest) (*types.ConstructionParseResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionPayloads Payloads is called with an array of operations and the response from
// /construction/metadata. It returns an unsigned transaction blob and a collection of payloads that
// must be signed by particular AccountIdentifiers using a certain SignatureType.
func (on OnlineNetwork) ConstructionPayloads(ctx context.Context, request *types.ConstructionPayloadsRequest) (*types.ConstructionPayloadsResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionPreprocess Preprocess is called prior to /construction/payloads to construct a
// request for any metadata that is needed for transaction construction given (i.e. account nonce).
func (on OnlineNetwork) ConstructionPreprocess(ctx context.Context, request *types.ConstructionPreprocessRequest) (*types.ConstructionPreprocessResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructionSubmit Submit a pre-signed transaction to the node. This call does not block on the
// transaction being included in a block. Rather, it returns immediately with an indication of
// whether or not the transaction was included in the mempool.
func (on OnlineNetwork) ConstructionSubmit(ctx context.Context, request *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	_ = "STUB: not implemented"
	return nil, nil
}
