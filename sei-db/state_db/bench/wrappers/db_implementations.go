package wrappers

import (
	"context"

	commonevm "github.com/sei-protocol/sei-chain/sei-db/common/keys"
	"github.com/sei-protocol/sei-chain/sei-db/config"
	flatkvConfig "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/flatkv/config"
	ssComposite "github.com/sei-protocol/sei-chain/sei-db/state_db/ss/composite"
)

const EVMStoreName = commonevm.EVMStoreKey

type DBType string

const (
	NoOp            DBType = "NoOp"
	MemIAVL         DBType = "MemIAVL"
	FlatKV          DBType = "FlatKV"
	CompositeDual   DBType = "CompositeDual"
	CompositeSplit  DBType = "CompositeSplit"
	CompositeCosmos DBType = "CompositeCosmos"

	SSComposite               DBType = "SSComposite"
	SSHistoricalOffload       DBType = "SSHistoricalOffload"
	CompositeDual_SSComposite DBType = "CompositeDual+SSComposite"
)

func DefaultBenchStateStoreConfig() *config.StateStoreConfig { _ = "STUB: not implemented"; return nil }

func newMemIAVLCommitStore(dbDir string) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}

func newFlatKVCommitStore(ctx context.Context, dbDir string, config *flatkvConfig.Config) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}

func newCompositeCommitStore(ctx context.Context, dbDir string, writeMode config.WriteMode) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}

func openSSComposite(dir string, cfg config.StateStoreConfig) (*ssComposite.CompositeStateStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSSCompositeStateStore(dbDir string, ssConfig *config.StateStoreConfig) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}

func newCombinedCompositeDualSSComposite(
	ctx context.Context,
	dbDir string,
	ssConfig *config.StateStoreConfig,
) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}

// NewDBImpl instantiates a new empty DBWrapper based on the given DBType.
func NewDBImpl(ctx context.Context, dbType DBType, dataDir string, dbConfig any) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}
