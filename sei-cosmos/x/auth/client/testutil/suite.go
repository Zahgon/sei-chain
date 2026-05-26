package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/sei-cosmos/client"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	_ = "STUB: not implemented"
	return nil
}

func (s *IntegrationTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

// Create a dummy account for testing purpose

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCLIValidateSignatures() { _ = "STUB: not implemented"; return }

// write  unsigned tx to file

func (s *IntegrationTestSuite) TestCLISignBatch() { _ = "STUB: not implemented"; return }

// sign-batch file - offline is set but account-number and sequence are not

// sign-batch file

// sign-batch file signature only

// Sign batch malformed tx file.

// Sign batch malformed tx file signature only.

func (s *IntegrationTestSuite) TestCLISignAminoJSON() { _ = "STUB: not implemented"; return }

// SIC! validators have same key names and same addresses as those registered in the keyring,
//      BUT the keys are different!

// query account info

/****  test signature-only  ****/

/****  test full output  ****/

// txCfg.UnmarshalSignatureJSON can't unmarshal a fragment of the signature, so we create this structure.

/****  test file output  ****/

/****  try to append to the previously signed transaction  ****/

/****  try to overwrite the previously signed transaction  ****/

// We can't sign with other address, because the bank send message supports only one signer for a simple
// account. Changing the file is too much hacking, because TxDecoder returns sdk.Tx, which doesn't
// provide functionality to check / manage `auth_info`.
// Cases with different keys are are covered in unit tests of `tx.Sign`.

/****  test flagAmino  ****/

func checkSignatures(require *require.Assertions, txCfg client.TxConfig, output []byte, pks ...cryptotypes.PubKey) {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestCLIQueryTxCmdByHash() { _ = "STUB: not implemented"; return }

// Send coins.

func (s *IntegrationTestSuite) TestCLIQueryTxCmdByEvents() { _ = "STUB: not implemented"; return }

// Send coins.

// Query the tx by hash to get the inner tx.

func (s *IntegrationTestSuite) TestCLIQueryTxsCmdByEvents() { _ = "STUB: not implemented"; return }

// Send coins.

// Query the tx by hash to get the inner tx.

func (s *IntegrationTestSuite) TestCLISendGenerateSignAndBroadcast() {
	_ = "STUB: not implemented"
	return
}

// Test generate sendTx with --gas=$amount

// Test generate sendTx, estimate gas

// Write the output to disk

// Test validate-signatures

// Test sign

// Does not work in offline mode

// But works offline if we set account number and sequence

// Sign transaction

// Write the output to disk

// validate Signature

// Ensure foo has right amount of funds

// Test broadcast

// Does not work in offline mode

// Broadcast correct transaction.

// Ensure destiny account state

// Ensure origin account state

func (s *IntegrationTestSuite) TestCLIMultisignInsufficientCosigners() {
	_ = "STUB: not implemented"
	return
}

// Fetch account and a multisig info

// Send coins from validator to multisig.

// Generate multisig transaction.

// Save tx to file

// Multisign, sign with one signature

// Save tx to file

func (s *IntegrationTestSuite) TestCLIEncode() { _ = "STUB: not implemented"; return }

// Encode

// Check that the transaction decodes as expected

func (s *IntegrationTestSuite) TestCLIMultisignSortSignatures() { _ = "STUB: not implemented"; return }

// Generate 2 accounts and a multisig.

// Generate dummy account which is not a part of multisig.

// Send coins from validator to multisig.

// Generate multisig transaction.

// Save tx to file

// Sign with account1

// Sign with account2

// Sign with dummy account

// Write the output to disk

func (s *IntegrationTestSuite) TestSignWithMultisig() { _ = "STUB: not implemented"; return }

// Generate a account for signing.

// Create an address that is not in the keyring, will be used to simulate `--multisig`

// Generate a transaction for testing --multisig with an address not in the keyring.

// Save multi tx to file

// Sign using multisig. We're signing a tx on behalf of the multisig address,
// even though the tx signer is NOT the multisig address. This is fine though,
// as the main point of this test is to test the `--multisig` flag with an address
// that is not in the keyring.

func (s *IntegrationTestSuite) TestCLIMultisign() { _ = "STUB: not implemented"; return }

// Generate 2 accounts and a multisig.

// Send coins from validator to multisig.

// Generate multisig transaction.

// Save tx to file

// Sign with account1

// Sign with account2

// Does not work in offline mode.

// Write the output to disk

func (s *IntegrationTestSuite) TestSignBatchMultisig() { _ = "STUB: not implemented"; return }

// Fetch 2 accounts and a multisig.

// Send coins from validator to multisig.

// Write the output to disk

// sign-batch file

// write sigs to file

// sign-batch file with account2

// write sigs to file2

// sign-batch file with multisig key name

// write sigs to file3

func (s *IntegrationTestSuite) TestMultisignBatch() { _ = "STUB: not implemented"; return }

// Fetch 2 accounts and a multisig.

// Send coins from validator to multisig.

// Write the output to disk

// sign-batch file

// write sigs to file

// sign-batch file with account2

// multisign the file

// sign-batch file with multisig key name

// write sigs to file

// Broadcast transactions.

func (s *IntegrationTestSuite) TestGetAccountCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetAccountsCmd() { _ = "STUB: not implemented"; return }

func TestGetBroadcastCommandOfflineFlag(t *testing.T) { _ = "STUB: not implemented"; return }

func TestGetBroadcastCommandWithoutOfflineFlag(t *testing.T) { _ = "STUB: not implemented"; return }

// Create new file with tx

func (s *IntegrationTestSuite) TestQueryParamsCmd() { _ = "STUB: not implemented"; return }

// TestTxWithoutPublicKey makes sure sending a proto tx message without the
// public key doesn't cause any error in the RPC layer (broadcast).
// See https://github.com/cosmos/cosmos-sdk/issues/7585 for more details.
func (s *IntegrationTestSuite) TestTxWithoutPublicKey() { _ = "STUB: not implemented"; return }

// Create a txBuilder with an unsigned tx.

// Set empty signature to set signer infos.

// Create a file with the unsigned tx.

// Sign the file with the unsignedTx.

// Remove the signerInfo's `public_key` field manually from the signedTx.
// Note: this method is only used for test purposes! In general, one should
// use txBuilder and TxEncoder/TxDecoder to manipulate txs.

// Re-encode the tx again, to another file.

// Broadcast tx, test that it shouldn't panic.

func (s *IntegrationTestSuite) TestSignWithMultiSignersAminoJSON() {
	_ = "STUB: not implemented"
	// test case:
	// Create a transaction with 2 messages which has to be signed with 2 different keys
	// Sign and append the signatures using the CLI with Amino signing mode.
	// Finally send the transaction to the blockchain. It should work.
	return
}

// Creating a tx with 2 msgs from 2 signers: val0 and val1.
// The validators need to sign with SIGN_MODE_LEGACY_AMINO_JSON,
// because DIRECT doesn't support multi signers via the CLI.
// Since we use amino, we don't need to pre-populate signer_infos.

// min required is 101892

// Write the unsigned tx into a file.

// Let val0 sign first the file with the unsignedTx.

// Then let val1 sign the file with signedByVal0.

// Now let's try to send this tx.

// Make sure the addr1's balance got funded.

func (s *IntegrationTestSuite) createBankMsg(val *network.Validator, toAddr sdk.AccAddress, amount sdk.Coins, extraFlags ...string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}
