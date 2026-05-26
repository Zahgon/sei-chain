package main

import (
	"math/rand"

	e2e "github.com/sei-protocol/sei-chain/sei-tendermint/test/e2e/pkg"
	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

var (
	// testnetCombinations defines global testnet options, where we generate a
	// separate testnet for each combination (Cartesian product) of options.
	testnetCombinations = map[string][]interface{}{
		"topology":      {"single", "quad", "large"},
		"initialHeight": {0, 1000},
		"initialState": {
			map[string]string{},
			map[string]string{"initial01": "a", "initial02": "b", "initial03": "c"},
		},
		"validators": {"genesis", "initchain"},
		"abci":       {"builtin", "outofprocess"},
	}

	// The following specify randomly chosen values for testnet nodes.
	nodeDatabases = weightedChoice{
		"goleveldb": 35,
		"badgerdb":  35,
		"boltdb":    15,
		"rocksdb":   10,
		"cleveldb":  5,
	}
	ABCIProtocols = weightedChoice{
		"tcp":  20,
		"grpc": 20,
		"unix": 10,
	}
	nodePrivvalProtocols = weightedChoice{
		"file": 50,
		"grpc": 20,
		"tcp":  20,
		"unix": 10,
	}
	nodeStateSyncs = weightedChoice{
		e2e.StateSyncDisabled: 10,
		e2e.StateSyncP2P:      45,
		e2e.StateSyncRPC:      45,
	}
	nodePersistIntervals  = uniformChoice{0, 1, 5}
	nodeSnapshotIntervals = uniformChoice{0, 5}
	nodeRetainBlocks      = uniformChoice{
		0,
		2 * int(e2e.EvidenceAgeHeight),
		4 * int(e2e.EvidenceAgeHeight),
	}
	nodePerturbations = probSetChoice{
		"disconnect": 0.1,
		"pause":      0.1,
		"kill":       0.1,
		"restart":    0.1,
	}

	// the following specify random chosen values for the entire testnet
	evidence   = uniformChoice{0, 1, 10}
	txSize     = uniformChoice{1024, 4096} // either 1kb or 4kb
	ipv6       = uniformChoice{false, true}
	keyType    = uniformChoice{types.ABCIPubKeyTypeEd25519}
	abciDelays = uniformChoice{"none", "small", "large"}
)

// Generate generates random testnets using the given RNG.
func Generate(r *rand.Rand, opts Options) ([]e2e.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	MinNetworkSize int
	MaxNetworkSize int
	NumGroups      int
	Directory      string
	Reverse        bool
}

// generateTestnet generates a single testnet with the given options.
func generateTestnet(r *rand.Rand, opt map[string]interface{}) (e2e.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(e2e.Manifest), nil
}

// FIXME Networks are kept small since large ones use too much CPU.

// First we generate seed nodes, starting at the initial height.

// Next, we generate validators. We make sure a BFT quorum of validators start
// at the initial height, and that we have two archive nodes. We also set up
// the initial validator set, and validator set updates for delayed nodes.

// Move validators to InitChain if specified.

// Finally, we generate random full nodes.

// We now set up peer discovery for nodes. Seed nodes are fully meshed with
// each other, while non-seed nodes either use a set of random seeds or a
// set of random peers that start before themselves.

// if the full node or validator is an ideal candidate, it is added as a light provider.
// There are at least two archive nodes so there should be at least two ideal candidates

// there are seeds, statesync is disabled, and it's
// either the first peer by the sort order, and
// (randomly half of the remaining peers use a seed
// node; otherwise, choose some remaining set of the
// peers.

// choose one of the seeds

// lastly, set up the light clients

// generateNode randomly generates a node, with some constraints to avoid
// generating invalid configurations. We do not set Seeds or PersistentPeers
// here, since we need to know the overall network topology and startup
// sequencing.
func generateNode(
	r *rand.Rand,
	manifest e2e.Manifest,
	mode e2e.Mode,
	startAt int64,
	forceArchive bool,
) *e2e.ManifestNode {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // test generator values are small non-negative ints
//nolint:gosec // test generator values are small non-negative ints
//nolint:gosec // test generator values are small non-negative ints

// avoid needing to blocsync more than five total blocks.

// If this node is forced to be an archive node, retain all blocks and
// enable state sync snapshotting.

// If a node which does not persist state also does not retain blocks, randomly
// choose to either persist state or retain all blocks.

// If either PersistInterval or SnapshotInterval are greater than RetainBlocks,
// expand the block retention time.

func generateLightNode(r *rand.Rand, startAt int64, providers []string) *e2e.ManifestNode {
	_ = "STUB: not implemented"
	return nil
}

func ptrUint64(i uint64) *uint64 { _ = "STUB: not implemented"; return nil }
