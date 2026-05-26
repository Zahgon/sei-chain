package littbuilder

import (
	"context"
	"log/slog"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/disktable/keymap"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/metrics"
)

// keymapBuilders contains builders for all supported keymap types.
var keymapBuilders = map[keymap.KeymapType]keymap.BuildKeymap{
	keymap.MemKeymapType:            keymap.NewMemKeymap,
	keymap.PebbleDBKeymapType:       keymap.NewPebbleDBKeymap,
	keymap.UnsafePebbleDBKeymapType: keymap.NewUnsafePebbleDBKeymap,
}

// cacheWeight is a function that calculates the weight of a cache entry.
func cacheWeight(key string, value []byte) uint64 { _ = "STUB: not implemented"; return 0 }

//nolint:gosec // lengths non-negative

// Look for a table's keymap directory in the provided segment paths.
func FindKeymapLocation(
	rootPaths []string,
	tableName string,
) (keymapDirectory string, keymapInitialized bool, keymapTypeFile *keymap.KeymapTypeFile, error error) {
	_ = "STUB: not implemented"
	return "", false, nil, nil
}

// buildKeymap creates a new keymap based on the configuration.
func buildKeymap(
	config *litt.Config,
	logger *slog.Logger,
	tableName string,
) (kmap keymap.Keymap, keymapPath string, keymapTypeFile *keymap.KeymapTypeFile, requiresReload bool, err error) {
	_ = "STUB: not implemented"
	return *new(keymap.Keymap), "", nil, false, nil
}

// The keymap has not been fully initialized. This is likely due to a crash during the keymap reloading process.

// No previous keymap exists. Either we are starting fresh or the keymap was deleted.

// by convention, always select the first path as the keymap directory

// create the keymap directory

// write the keymap type file

// A previous keymap exists. Check if the keymap type has changed.

// The previously used keymap type is different from the one in the configuration.

// delete the old keymap

// write the new keymap type file

// If the keymap does not need to be reloaded, then it is already fully initialized.

//nolint:gosec // path within keymap directory

// buildTable creates a new table based on the configuration.
func buildTable(
	config *litt.Config,
	logger *slog.Logger,
	name string,
	metrics *metrics.LittDBMetrics) (litt.ManagedTable, error) {
	_ = "STUB: not implemented"
	return *new(litt.ManagedTable), nil
}

// buildLogger returns the configured logger or slog.Default() if none was provided.
func buildLogger(config *litt.Config) *slog.Logger { _ = "STUB: not implemented"; return nil }

// buildMetrics creates a new metrics object backed by the global OTel
// MeterProvider. When MetricsEnabled is true, this configures the global
// provider with a Prometheus exporter and starts an HTTP server on
// MetricsPort that serves /metrics. The returned shutdown function flushes
// the provider; it is the responsibility of the caller to invoke it during
// teardown.
func buildMetrics(config *litt.Config, logger *slog.Logger) (*metrics.LittDBMetrics, func(context.Context) error) {
	_ = "STUB: not implemented"
	return nil, nil
}
