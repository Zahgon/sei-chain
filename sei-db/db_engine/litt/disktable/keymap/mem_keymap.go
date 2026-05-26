package keymap

import (
	"log/slog"
	"sync"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/litt/types"
)

var _ Keymap = &memKeymap{}

// An in-memory keymap implementation. When a table using a memKeymap is restarted, it loads all keys from
// the segment files.  Methods on this struct are goroutine safe.
//
// - potentially high memory usage for large keymaps
// - potentially slow startup time for large keymaps
// - very fast after startup
type memKeymap struct {
	logger *slog.Logger
	data   map[string]types.Address
	// if true, then return an error if an update would overwrite an existing key
	doubleWriteProtection bool
	lock                  sync.RWMutex
}

var _ BuildKeymap = NewMemKeymap

// NewMemKeymap creates a new in-memory keymap.
func NewMemKeymap(logger *slog.Logger,
	_ string,
	doubleWriteProtection bool) (kmap Keymap, requiresReload bool, err error) {
	_ = "STUB: not implemented"
	return *new(Keymap), false, nil
}

func (m *memKeymap) Put(keys []*types.ScopedKey) error { _ = "STUB: not implemented"; return nil }

func (m *memKeymap) Get(key []byte) (types.Address, bool, error) {
	_ = "STUB: not implemented"
	return *new(types.Address), false, nil
}

func (m *memKeymap) Delete(keys []*types.ScopedKey) error { _ = "STUB: not implemented"; return nil }

func (m *memKeymap) Stop() error {
	_ = "STUB: not implemented"
	// nothing to do here
	return nil
}

func (m *memKeymap) Destroy() error { _ = "STUB: not implemented"; return nil }
