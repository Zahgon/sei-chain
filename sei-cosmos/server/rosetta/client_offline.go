package rosetta

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// ---------- cosmos-rosetta-gateway.types.NetworkInformationProvider implementation ------------ //

func (c *Client) OperationStatuses() []*types.OperationStatus {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Version() string { _ = "STUB: not implemented"; return "" }

func (c *Client) SupportedOperations() []string { _ = "STUB: not implemented"; return nil }

// ---------- cosmos-rosetta-gateway.types.OfflineClient implementation ------------ //

func (c *Client) SignedTx(_ context.Context, txBytes []byte, signatures []*types.Signature) (signedTxBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) ConstructionPayload(_ context.Context, request *types.ConstructionPayloadsRequest) (resp *types.ConstructionPayloadsResponse, err error) {
	_ = "STUB: not implemented"
	// check if there is at least one operation
	return nil, nil
}

func (c *Client) PreprocessOperationsToOptions(_ context.Context, req *types.ConstructionPreprocessRequest) (response *types.ConstructionPreprocessResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// now we need to parse the operations to cosmos sdk messages

// get the signers

// get the metadata request information

// prepare the options to return

func (c *Client) AccountIdentifierFromPublicKey(pubKey *types.PublicKey) (*types.AccountIdentifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
