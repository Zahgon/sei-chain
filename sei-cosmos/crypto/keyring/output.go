package keyring

import (
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

// TODO: Move this file to client/keys
// Use protobuf interface marshaler rather then generic JSON

// KeyOutput defines a structure wrapping around an Info object used for output
// functionality.
type KeyOutput struct {
	Name       string `json:"name" yaml:"name"`
	Type       string `json:"type" yaml:"type"`
	Address    string `json:"address" yaml:"address"`
	EvmAddress string `json:"evm_address" yaml:"evm_address"`
	PubKey     string `json:"pubkey" yaml:"pubkey"`
	Mnemonic   string `json:"mnemonic,omitempty" yaml:"mnemonic"`
}

// NewKeyOutput creates a default KeyOutput instance without Mnemonic, Threshold and PubKeys
func NewKeyOutput(name string, keyType KeyType, a sdk.Address, pk cryptotypes.PubKey) (KeyOutput, error) {
	_ = "STUB: not implemented"
	return *new(KeyOutput), nil
}

// MkConsKeyOutput create a KeyOutput in with "cons" Bech32 prefixes.
func MkConsKeyOutput(keyInfo Info) (KeyOutput, error) {
	_ = "STUB: not implemented"
	return *new(KeyOutput), nil
}

// MkValKeyOutput create a KeyOutput in with "val" Bech32 prefixes.
func MkValKeyOutput(keyInfo Info) (KeyOutput, error) {
	_ = "STUB: not implemented"
	return *new(KeyOutput), nil
}

// MkAccKeyOutput create a KeyOutput in with "acc" Bech32 prefixes. If the
// public key is a multisig public key, then the threshold and constituent
// public keys will be added.
func MkAccKeyOutput(keyInfo Info) (KeyOutput, error) {
	_ = "STUB: not implemented"
	return *new(KeyOutput), nil
}

// MkAccKeysOutput returns a slice of KeyOutput objects, each with the "acc"
// Bech32 prefixes, given a slice of Info objects. It returns an error if any
// call to MkKeyOutput fails.
func MkAccKeysOutput(infos []Info) ([]KeyOutput, error) { _ = "STUB: not implemented"; return nil, nil }

// We only support getting evm-addr if the algo type is secp256k1 (which it should be, though there
// may be some legacy keys with sr25519)

func PopulateEvmAddrIfApplicable(info Info, o KeyOutput) (KeyOutput, error) {
	_ = "STUB: not implemented"
	return *new(KeyOutput), nil
}

// Only works with secp256k1 algo
