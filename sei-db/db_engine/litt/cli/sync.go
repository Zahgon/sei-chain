package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/urfave/cli/v2"
)

func syncCommand(ctx *cli.Context) error { _ = "STUB: not implemented"; return nil }

// A utility that periodically transfers data from a local database to a remote backup using rsync.
type syncEngine struct {
	ctx                 context.Context
	cancel              context.CancelFunc
	logger              *slog.Logger
	sources             []string
	destinations        []string
	user                string
	host                string
	port                uint64
	keyPath             string
	knownHostsFile      string
	deleteAfterTransfer bool
	fsync               bool
	threads             uint64
	throttleMB          float64
	period              time.Duration
	maxAgeSeconds       uint64
	remoteLittBinary    string
	verbose             bool
}

// newSyncEngine creates a new syncEngine instance with the provided parameters.
func newSyncEngine(
	ctx context.Context,
	logger *slog.Logger,
	sources []string,
	destinations []string,
	user string,
	host string,
	port uint64,
	keyPath string,
	knownHostsFile string,
	deleteAfterTransfer bool,
	fsync bool,
	threads uint64,
	throttleMB float64,
	period time.Duration,
	maxAgeSeconds uint64,
	remoteLittBinary string,
	verbose bool,
) *syncEngine {
	_ = "STUB: not implemented"
	return nil
}

// run the sync engine. This method blocks until the context is cancelled or an unrecoverable error occurs.
func (s *syncEngine) run() error {
	_ = "STUB: not implemented"

	// Create a channel to listen for OS signals
	return nil
}

// Wait for signal

// Cancel the context when signal is received

// syncLoop is the main loop of the sync engine. It runs indefinitely until the context is cancelled.
func (s *syncEngine) syncLoop() { _ = "STUB: not implemented"; return }

func (s *syncEngine) sync() { _ = "STUB: not implemented"; return }

// Stop stops the sync engine by cancelling the context.
func (s *syncEngine) Stop() { _ = "STUB: not implemented"; return }
