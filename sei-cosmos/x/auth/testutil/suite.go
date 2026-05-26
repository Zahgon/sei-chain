package testutil

import (
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	signingtypes "github.com/sei-protocol/sei-chain/sei-cosmos/types/tx/signing"
)

// TxConfigTestSuite provides a test suite that can be used to test that a TxConfig implementation is correct.
type TxConfigTestSuite struct {
	suite.Suite
	TxConfig client.TxConfig
}

// NewTxConfigTestSuite returns a new TxConfigTestSuite with the provided TxConfig implementation
func NewTxConfigTestSuite(txConfig client.TxConfig) *TxConfigTestSuite {
	_ = "STUB: not implemented"
	return nil
}

func (s *TxConfigTestSuite) TestTxBuilderGetTx() { _ = "STUB: not implemented"; return }

func (s *TxConfigTestSuite) TestTxBuilderSetFeeAmount() { _ = "STUB: not implemented"; return }

func (s *TxConfigTestSuite) TestTxBuilderSetGasLimit() { _ = "STUB: not implemented"; return }

func (s *TxConfigTestSuite) TestTxBuilderSetMemo() { _ = "STUB: not implemented"; return }

func (s *TxConfigTestSuite) TestTxBuilderSetMsgs() { _ = "STUB: not implemented"; return }

// should fail because of no signatures

func (s *TxConfigTestSuite) TestTxBuilderSetSignatures() { _ = "STUB: not implemented"; return }

// set test msg

// check that validation fails

// set SignatureV2 without actual signature bytes
// Arbitrary account sequence

// Arbitrary account sequence

// fail validation without required signers

// sign transaction

// set signature

func sigEquals(sig1, sig2 signingtypes.SignatureV2) bool { _ = "STUB: not implemented"; return false }

func sigDataEquals(data1, data2 signingtypes.SignatureData) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *TxConfigTestSuite) TestTxEncodeDecode() { _ = "STUB: not implemented"; return }

func (s *TxConfigTestSuite) TestWrapTxBuilder() { _ = "STUB: not implemented"; return }
