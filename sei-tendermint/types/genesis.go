package types

import (
	"encoding/json"
	"time"

	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"
	"github.com/sei-protocol/sei-chain/sei-tendermint/crypto"
	tmbytes "github.com/sei-protocol/sei-chain/sei-tendermint/libs/bytes"
)

const (
	// MaxChainIDLen is a maximum length of the chain ID.
	MaxChainIDLen = 50
)

//------------------------------------------------------------
// core types for a genesis definition
// NOTE: any changes to the genesis definition should
// be reflected in the documentation:
// docs/tendermint-core/using-tendermint.md

// GenesisValidator is an initial validator.
type GenesisValidator struct {
	Address Address
	PubKey  crypto.PubKey
	Power   int64
	Name    string
}

type genesisValidatorJSON struct {
	Address Address         `json:"address"`
	PubKey  json.RawMessage `json:"pub_key"`
	Power   int64           `json:"power,string"`
	Name    string          `json:"name"`
}

func (g GenesisValidator) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *GenesisValidator) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// GenesisDoc defines the initial conditions for a tendermint blockchain, in particular its validator set.
type GenesisDoc struct {
	GenesisTime     time.Time          `json:"genesis_time"`
	ChainID         string             `json:"chain_id"`
	InitialHeight   int64              `json:"initial_height,string"`
	ConsensusParams *ConsensusParams   `json:"consensus_params,omitempty"`
	Validators      []GenesisValidator `json:"validators,omitempty"`
	AppHash         tmbytes.HexBytes   `json:"app_hash"`
	AppState        json.RawMessage    `json:"app_state,omitempty"`
}

func (genDoc *GenesisDoc) ToRequestInitChain() *abci.RequestInitChain {
	_ = "STUB: not implemented"
	return nil
}

// SaveAs is a utility method for saving GenensisDoc as a JSON file.
func (genDoc *GenesisDoc) SaveAs(file string) error { _ = "STUB: not implemented"; return nil }

// nolint:gosec

func (genDoc *GenesisDoc) ValidatorSet() *ValidatorSet { _ = "STUB: not implemented"; return nil }

func (genDoc *GenesisDoc) ValidatorUpdates() []abci.ValidatorUpdate {
	_ = "STUB: not implemented"
	return nil
}

// ValidatorHash returns the hash of the validator set contained in the GenesisDoc
func (genDoc *GenesisDoc) ValidatorHash() []byte { _ = "STUB: not implemented"; return nil }

// ValidateAndComplete checks that all necessary fields are present
// and fills in defaults for optional fields left empty
func (genDoc *GenesisDoc) ValidateAndComplete() error { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------
// Make genesis state from file

// GenesisDocFromJSON unmarshalls JSON data into a GenesisDoc.
func GenesisDocFromJSON(jsonBlob []byte) (*GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenesisDocFromFile reads JSON data from a file and unmarshalls it into a GenesisDoc.
func GenesisDocFromFile(genDocFile string) (*GenesisDoc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
