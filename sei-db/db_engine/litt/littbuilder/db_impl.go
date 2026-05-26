package littbuilder

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/metrics"
)

var _ litt.DB = &db{}

// TableBuilderFunc is a function that creates a new table.
type TableBuilderFunc func(
	ctx context.Context,
	logger *slog.Logger,
	name string,
	metrics *metrics.LittDBMetrics) (litt.ManagedTable, error)

// db is an implementation of DB.
type db struct {
	ctx    context.Context
	logger *slog.Logger

	// A function that returns the current time.
	clock func() time.Time

	// The default time-to-live for new tables. Once created, the TTL for a table can be changed.
	ttl time.Duration

	// The period between garbage collection runs.
	gcPeriod time.Duration

	// A function that creates new tables.
	tableBuilder TableBuilderFunc

	// A map of all tables in the database.
	tables map[string]litt.ManagedTable

	// Protects access to tables and ttl.
	lock sync.Mutex

	// True if the database has been stopped.
	stopped atomic.Bool

	// Metrics for the database.
	metrics *metrics.LittDBMetrics

	// Shuts down the OTel MeterProvider configured by buildMetrics. nil if metrics are disabled.
	metricsShutdown func(context.Context) error

	// A function that releases file locks.
	releaseLocks func()

	// Set to true when the database is closed.
	closed bool
}

// NewDB creates a new DB instance. After this method is called, the config object should not be modified.
func NewDB(config *litt.Config) (litt.DB, error) {
	_ = "STUB: not implemented"
	return *new(litt.DB), nil
}

// NewDBUnsafe creates a new DB instance with a custom table builder. This is intended for unit test use,
// and should not be considered a stable API.
func NewDBUnsafe(config *litt.Config, tableBuilder TableBuilderFunc) (litt.DB, error) {
	_ = "STUB: not implemented"
	return *new(litt.DB), nil
}

func (d *db) KeyCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *db) Size() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *db) lockFreeSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *db) GetTable(name string) (litt.Table, error) {
	_ = "STUB: not implemented"
	return *new(litt.Table), nil
}

func (d *db) DropTable(name string) error { _ = "STUB: not implemented"; return nil }

// Table does not exist, nothing to do.

func (d *db) Close() error { _ = "STUB: not implemented"; return nil }

func (d *db) closeUnsafe() error {
	_ = "STUB: not implemented"

	// closing more than once is a no-op
	return nil
}

func (d *db) Destroy() error { _ = "STUB: not implemented"; return nil }

// gatherMetrics is a method that periodically collects metrics.
func (d *db) gatherMetrics(interval time.Duration) { _ = "STUB: not implemented"; return }
