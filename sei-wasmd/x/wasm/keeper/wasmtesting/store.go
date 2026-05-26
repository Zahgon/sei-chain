package wasmtesting

import (
	storetypes "github.com/sei-protocol/sei-chain/sei-cosmos/store/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// MockCommitMultiStore mock with a CacheMultiStore to capture commits
type MockCommitMultiStore struct {
	sdk.CommitMultiStore
	Committed []bool
}

func (m *MockCommitMultiStore) CacheMultiStore() storetypes.CacheMultiStore {
	_ = "STUB: not implemented"
	return *new(storetypes.CacheMultiStore)
}

type mockCMS struct {
	sdk.CommitMultiStore
	committed *bool
}

func (m *mockCMS) Close() { _ = "STUB: not implemented"; return }

func (m *mockCMS) Write() { _ = "STUB: not implemented"; return }
