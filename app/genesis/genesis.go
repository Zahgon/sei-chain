package genesis

import (
	"embed"

	"github.com/sei-protocol/sei-chain/sei-tendermint/types"
)

//go:embed chains/*.json
var embeddedGenesisFS embed.FS

func WellKnownChainIDs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func IsWellKnown(chainID string) bool { _ = "STUB: not implemented"; return false }

func EmbeddedGenesis(chainID string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func EmbeddedGenesisDoc(chainID string) (*types.GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
