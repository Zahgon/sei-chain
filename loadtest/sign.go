package main

import (
	"sync"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

type AccountInfo struct {
	Address  string `json:"address"`
	Mnemonic string `json:"mnemonic"`
}

type SignerInfo struct {
	AccountNumber  uint64
	SequenceNumber uint64
	mutex          *sync.Mutex
}

func NewSignerInfo(accountNumber uint64, sequenceNumber uint64) *SignerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (si *SignerInfo) IncrementAccountNumber() { _ = "STUB: not implemented"; return }

type SignerClient struct {
	CachedAccountSeqNum *sync.Map
	CachedAccountKey    *sync.Map
	NodeURI             string
}

func NewSignerClient(nodeURI string) *SignerClient { _ = "STUB: not implemented"; return nil }

func (sc *SignerClient) GetTestAccountsKeys(maxAccounts int) []cryptotypes.PrivKey {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SignerClient) GetAdminAccountKeyPath() string { _ = "STUB: not implemented"; return "" }

func (sc *SignerClient) GetAdminKey() cryptotypes.PrivKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PrivKey)
}

func (sc *SignerClient) GetKey(accountID, backend, accountKeyFilePath string) cryptotypes.PrivKey {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PrivKey)
}

// Cache this so we don't need to regenerate it

func (sc *SignerClient) GetValKeys() ([]cryptotypes.PrivKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we dont expect subdirectories, so we can just handle files

//nolint:gosec // used for testing only

func (sc *SignerClient) SignTx(chainID string, txBuilder *client.TxBuilder, privKey cryptotypes.PrivKey, seqDelta uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SignerClient) GetAccountNumberSequenceNumber(privKey cryptotypes.PrivKey) SignerInfo {
	_ = "STUB: not implemented"
	return *new(SignerInfo)
}

// retry once after 5 seconds
