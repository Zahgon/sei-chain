package sstest

import (
	"math/rand"
	"testing"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
)

// StorageBenchSuite defines a reusable benchmark suite for all storage backends.
type StorageBenchSuite struct {
	BenchBackendName string
	NewDB            func(dir string) (types.StateStore, error)
}

var rng = rand.New(rand.NewSource(567320))

func (s *StorageBenchSuite) BenchmarkGet(b *testing.B) { _ = "STUB: not implemented"; return }

func (s *StorageBenchSuite) BenchmarkApplyChangeset(b *testing.B) {
	_ = "STUB: not implemented"
	return
}

func (s *StorageBenchSuite) BenchmarkIterate(b *testing.B) { _ = "STUB: not implemented"; return }
