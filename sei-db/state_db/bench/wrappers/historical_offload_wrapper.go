package wrappers

import (
	"context"
	"sync/atomic"

	"github.com/sei-protocol/sei-chain/sei-db/common/metrics"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
	scTypes "github.com/sei-protocol/sei-chain/sei-db/state_db/sc/types"
	"github.com/sei-protocol/sei-chain/sei-db/state_db/ss/offload"
)

var _ DBWrapper = (*historicalOffloadWrapper)(nil)

type HistoricalOffloadConfig struct {
	Provider string
	Kafka    *KafkaHistoricalOffloadConfig
}

type KafkaHistoricalOffloadConfig struct {
	Brokers        []string
	Topic          string
	ClientID       string
	Region         string
	Async          *bool
	RequiredAcks   string
	Compression    string
	BatchSize      int
	BatchTimeoutMS int
	BatchBytes     int
	TLSEnabled     bool
	SASLMechanism  string
}

type historicalOffloadWrapper struct {
	stream  offload.Stream
	version atomic.Int64
}

func (c *HistoricalOffloadConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *KafkaHistoricalOffloadConfig) applyDefaults() { _ = "STUB: not implemented"; return }

func (c *KafkaHistoricalOffloadConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (c *KafkaHistoricalOffloadConfig) asyncValue() bool { _ = "STUB: not implemented"; return false }

func newHistoricalOffloadStream(cfg *HistoricalOffloadConfig) (offload.Stream, error) {
	_ = "STUB: not implemented"
	return *new(offload.Stream), nil
}

func newSSHistoricalOffloadStateStore(_ context.Context, dbDir string, cfg *HistoricalOffloadConfig) (DBWrapper, error) {
	_ = "STUB: not implemented"
	return *new(DBWrapper), nil
}

func NewHistoricalOffloadWrapper(stream offload.Stream) DBWrapper {
	_ = "STUB: not implemented"
	return *new(DBWrapper)
}

func (h *historicalOffloadWrapper) ApplyChangeSets(entry *proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *historicalOffloadWrapper) Read(_ []byte) (data []byte, found bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (h *historicalOffloadWrapper) Commit() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *historicalOffloadWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (h *historicalOffloadWrapper) Version() int64 { _ = "STUB: not implemented"; return 0 }

func (h *historicalOffloadWrapper) LoadVersion(_ int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *historicalOffloadWrapper) Importer(_ int64) (scTypes.Importer, error) {
	_ = "STUB: not implemented"
	return *new(scTypes.Importer), nil
}

func (h *historicalOffloadWrapper) GetPhaseTimer() *metrics.PhaseTimer {
	_ = "STUB: not implemented"
	return nil
}
