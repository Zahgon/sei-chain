package keymap

import (
	"log/slog"
	"sync/atomic"

	"github.com/cockroachdb/pebble/v2"
	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/types"
)

var _ Keymap = &PebbleDBKeymap{}

// PebbleDBKeymap is a keymap that uses PebbleDB as the underlying storage. Methods on this struct are goroutine safe.
type PebbleDBKeymap struct {
	logger *slog.Logger
	db     *pebble.DB
	// if true, then return an error if an update would overwrite an existing key
	doubleWriteProtection bool
	keymapPath            string
	alive                 atomic.Bool
	// This is a "test mode only" flag. Should be true in production use cases or anywhere that data consistency
	// is critical. Unit tests write lots of little values, and syncing each one is slow, so it may be desirable
	// to set this to false in some tests.
	syncWrites bool
}

var _ BuildKeymap = NewPebbleDBKeymap

// NewPebbleDBKeymap creates a new PebbleDBKeymap instance.
func NewPebbleDBKeymap(
	logger *slog.Logger,
	keymapPath string,
	doubleWriteProtection bool) (kmap Keymap, requiresReload bool, err error) {
	_ = "STUB: not implemented"
	return *new(Keymap), false, nil
}

// NewUnsafePebbleDBKeymap creates a new PebbleDBKeymap instance. It does not use sync writes. This makes it faster,
// but unsafe if data consistency is critical (i.e. production use cases).
func NewUnsafePebbleDBKeymap(
	logger *slog.Logger,
	keymapPath string,
	doubleWriteProtection bool) (kmap Keymap, requiresReload bool, err error) {
	_ = "STUB: not implemented"
	return *new(Keymap), false, nil
}

// newPebbleDBKeymap creates a new PebbleDBKeymap instance.
func newPebbleDBKeymap(
	logger *slog.Logger,
	keymapPath string,
	doubleWriteProtection bool,
	syncWrites bool) (kmap *PebbleDBKeymap, requiresReload bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (p *PebbleDBKeymap) writeOptions() *pebble.WriteOptions { _ = "STUB: not implemented"; return nil }

func (p *PebbleDBKeymap) Put(keys []*types.ScopedKey) error { _ = "STUB: not implemented"; return nil }

func (p *PebbleDBKeymap) Get(key []byte) (types.Address, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Address), false, nil
}

// Clone the bytes before closing, since the slice is only valid until closer.Close().

func (p *PebbleDBKeymap) Delete(keys []*types.ScopedKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PebbleDBKeymap) Stop() error { _ = "STUB: not implemented"; return nil }

func (p *PebbleDBKeymap) Destroy() error { _ = "STUB: not implemented"; return nil }
