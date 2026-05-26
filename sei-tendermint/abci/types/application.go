package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Application is an interface that enables any finite, deterministic state machine
// to be driven by a blockchain-based replication engine via the ABCI.
//
//go:generate ../../scripts/mockery_generate.sh Application
type Application interface {
	// Info/Query Connection
	Info(context.Context, *RequestInfo) (*ResponseInfo, error)    // Return application info
	Query(context.Context, *RequestQuery) (*ResponseQuery, error) // Query for state
	GetValidators() []ValidatorUpdate

	// Mempool Connection
	CheckTx(context.Context, *RequestCheckTxV2) *ResponseCheckTxV2                                      // Validate a tx for the mempool
	GetTxPriorityHint(context.Context, *RequestGetTxPriorityHintV2) (*ResponseGetTxPriorityHint, error) // Get tx priority before checkTx
	EvmNonce(common.Address) uint64
	EvmBalance(common.Address, []byte) *big.Int

	// Consensus Connection
	InitChain(context.Context, *RequestInitChain) (*ResponseInitChain, error) // Initialize blockchain w validators/other info from TendermintCore
	ProcessProposal(context.Context, *RequestProcessProposal) (*ResponseProcessProposal, error)
	// Commit the state and return the application Merkle root hash
	Commit(context.Context) (*ResponseCommit, error)
	// Deliver the decided block with its txs to the Application
	FinalizeBlock(context.Context, *RequestFinalizeBlock) (*ResponseFinalizeBlock, error)

	// State Sync Connection
	ListSnapshots(context.Context, *RequestListSnapshots) (*ResponseListSnapshots, error)                // List available snapshots
	OfferSnapshot(context.Context, *RequestOfferSnapshot) (*ResponseOfferSnapshot, error)                // Offer a snapshot to the application
	LoadSnapshotChunk(context.Context, *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error)    // Load a snapshot chunk
	ApplySnapshotChunk(context.Context, *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) // Apply a shapshot chunk
}

//-------------------------------------------------------
// BaseApplication is a base form of Application

var _ Application = BaseApplication{}

type BaseApplication struct{}

func (BaseApplication) Info(_ context.Context, req *RequestInfo) (*ResponseInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) GetValidators() []ValidatorUpdate { _ = "STUB: not implemented"; return nil }

func (BaseApplication) CheckTx(_ context.Context, req *RequestCheckTxV2) *ResponseCheckTxV2 {
	_ = "STUB: not implemented"
	return nil
}

func (BaseApplication) Commit(_ context.Context) (*ResponseCommit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) Query(_ context.Context, req *RequestQuery) (*ResponseQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) InitChain(_ context.Context, req *RequestInitChain) (*ResponseInitChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ListSnapshots(_ context.Context, req *RequestListSnapshots) (*ResponseListSnapshots, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) OfferSnapshot(_ context.Context, req *RequestOfferSnapshot) (*ResponseOfferSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) LoadSnapshotChunk(_ context.Context, _ *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ApplySnapshotChunk(_ context.Context, req *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) ProcessProposal(_ context.Context, req *RequestProcessProposal) (*ResponseProcessProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) GetTxPriorityHint(context.Context, *RequestGetTxPriorityHintV2) (*ResponseGetTxPriorityHint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (BaseApplication) EvmNonce(common.Address) uint64 { _ = "STUB: not implemented"; return 0 }

func (BaseApplication) EvmBalance(common.Address, []byte) *big.Int {
	_ = "STUB: not implemented"
	return nil
}

func (BaseApplication) FinalizeBlock(_ context.Context, req *RequestFinalizeBlock) (*ResponseFinalizeBlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
