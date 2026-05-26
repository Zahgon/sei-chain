package evmrpc

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/crypto"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
)

var errNoProofCapableQueryableKVStore = errors.New("cannot find a proof-capable queryable KV store")

type StateAPI struct {
	tmClient       client.LocalClient
	keeper         *keeper.Keeper
	ctxProvider    func(int64) sdk.Context
	connectionType ConnectionType
	watermarks     *WatermarkManager
}

func NewStateAPI(tmClient client.LocalClient, k *keeper.Keeper, ctxProvider func(int64) sdk.Context, connectionType ConnectionType, watermarks *WatermarkManager) *StateAPI {
	_ = "STUB: not implemented"
	return nil
}

func (a *StateAPI) GetBalance(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (result *hexutil.Big, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *StateAPI) GetCode(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (result hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

func (a *StateAPI) GetStorageAt(ctx context.Context, address common.Address, hexKey string, blockNrOrHash rpc.BlockNumberOrHash) (result hexutil.Bytes, returnErr error) {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes), nil
}

// Result structs for GetProof
// This differs from go-ethereum AccountResult in two ways:
// 1. Proof object is an iavl proof, not a trie proof
// 2. Per-account fields are excluded because there is no per-account root
type ProofResult struct {
	Address      common.Address     `json:"address"`
	HexValues    []string           `json:"hexValues"`
	StorageProof []*crypto.ProofOps `json:"storageProof"`
}

func (a *StateAPI) GetProof(ctx context.Context, address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (result *ProofResult, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findQueryableKVStore unwraps known KVStore wrappers until it reaches a types.Queryable
// (classic IAVL, store/v2 memiavl commitment, or future proof-capable roots).
// Go only allows `x := s.(type)` inside a type switch, not before it. Nil parents are
// handled by the `s == nil` check on the next iteration; nil *Store receivers are
// guarded in each pointer case so we never call methods on nil.
func findQueryableKVStore(s sdk.KVStore) (storetypes.Queryable, error) {
	_ = "STUB: not implemented"
	return *new(storetypes.Queryable), nil
}

func (a *StateAPI) GetNonce(ctx context.Context, address common.Address) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// decodeHash parses a hex-encoded 32-byte hash. The input may optionally
// be prefixed by 0x and can have a byte length up to 32.
func decodeHash(s string) (h common.Hash, inputLength int, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), 0, nil
}
