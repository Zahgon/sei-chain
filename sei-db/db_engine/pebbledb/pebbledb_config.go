package pebbledb

import (
	"time"
)

// Configuration for the PebbleDB database.
type PebbleDBConfig struct {
	// The directory to store the database files. This has no default value and must be provided.
	DataDir string
	// Whether to enable pebble-internal metrics.
	EnableMetrics bool
	// How often to scrape pebble-internal metrics.
	MetricsScrapeInterval time.Duration
}

// Default configuration for the PebbleDB database.
func DefaultConfig() PebbleDBConfig { _ = "STUB: not implemented"; return *new(PebbleDBConfig) }

// Validates the configuration (basic sanity checks).
func (c *PebbleDBConfig) Validate() error { _ = "STUB: not implemented"; return nil }
