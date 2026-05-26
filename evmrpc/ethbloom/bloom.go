package ethbloom

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
)

var bitMasks = [8]uint8{1, 2, 4, 8, 16, 32, 64, 128}

// BloomIndexes represents the bit indexes inside the bloom filter that belong
// to some key.
type BloomIndexes [3]uint

func calcBloomIndexes(b []byte) BloomIndexes { _ = "STUB: not implemented"; return *new(BloomIndexes) }

// EncodeFilters builds bloom-index slices from filter criteria.
// Result semantics: AND on outer level, OR on mid level, AND on inner level (all 3 bits).
func EncodeFilters(addresses []common.Address, topics [][]common.Hash) (res [][]BloomIndexes) {
	_ = "STUB: not implemented"
	return nil
}

// MatchFilters returns true when bloom matches all filter groups.
func MatchFilters(bloom ethtypes.Bloom, filterGroups [][]BloomIndexes) bool {
	_ = "STUB: not implemented"
	return false
}

func matchFilter(bloom ethtypes.Bloom, filter []BloomIndexes) bool {
	_ = "STUB: not implemented"
	return false
}

func matchBloomIndexes(bloom ethtypes.Bloom, idx BloomIndexes) bool {
	_ = "STUB: not implemented"
	return false
}

// MatchesCriteria checks if a log matches the filter criteria (exact match).
func MatchesCriteria(log *ethtypes.Log, crit filters.FilterCriteria) bool {
	_ = "STUB: not implemented"
	return false
}
