package rosetta

import (
	tmcoretypes "github.com/sei-protocol/sei-chain/sei-tendermint/rpc/coretypes"
	tmtypes "github.com/sei-protocol/sei-chain/sei-tendermint/types"

	crgtypes "github.com/sei-protocol/sei-chain/sei-cosmos/server/rosetta/lib/types"

	rosettatypes "github.com/coinbase/rosetta-sdk-go/types"
	abci "github.com/sei-protocol/sei-chain/sei-tendermint/abci/types"

	sdkclient "github.com/sei-protocol/sei-chain/sei-cosmos/client"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	authsigning "github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/signing"
)

// Converter is a utility that can be used to convert
// back and forth from rosetta to sdk and tendermint types
// IMPORTANT NOTES:
//   - IT SHOULD BE USED ONLY TO DEAL WITH THINGS
//     IN A STATELESS WAY! IT SHOULD NEVER INTERACT DIRECTLY
//     WITH TENDERMINT RPC AND COSMOS GRPC
//
// - IT SHOULD RETURN cosmos rosetta gateway error types!
type Converter interface {
	// ToSDK exposes the methods that convert
	// rosetta types to cosmos sdk and tendermint types
	ToSDK() ToSDKConverter
	// ToRosetta exposes the methods that convert
	// sdk and tendermint types to rosetta types
	ToRosetta() ToRosettaConverter
}

// ToRosettaConverter is an interface that exposes
// all the functions used to convert sdk and
// tendermint types to rosetta known types
type ToRosettaConverter interface {
	// BlockResponse returns a block response given a result block
	BlockResponse(block *tmcoretypes.ResultBlock) crgtypes.BlockResponse
	// BeginBlockToTx converts the given begin block hash to rosetta transaction hash
	BeginBlockTxHash(blockHash []byte) string
	// EndBlockTxHash converts the given endblock hash to rosetta transaction hash
	EndBlockTxHash(blockHash []byte) string
	// Amounts converts sdk.Coins to rosetta.Amounts
	Amounts(ownedCoins []sdk.Coin, availableCoins sdk.Coins) []*rosettatypes.Amount
	// Ops converts an sdk.Msg to rosetta operations
	Ops(status string, msg sdk.Msg) ([]*rosettatypes.Operation, error)
	// OpsAndSigners takes raw transaction bytes and returns rosetta operations and the expected signers
	OpsAndSigners(txBytes []byte) (ops []*rosettatypes.Operation, signers []*rosettatypes.AccountIdentifier, err error)
	// Meta converts an sdk.Msg to rosetta metadata
	Meta(msg sdk.Msg) (meta map[string]interface{}, err error)
	// SignerData returns account signing data from a queried any account
	SignerData(anyAccount *codectypes.Any) (*SignerData, error)
	// SigningComponents returns rosetta's components required to build a signable transaction
	SigningComponents(tx authsigning.Tx, metadata *ConstructionMetadata, rosPubKeys []*rosettatypes.PublicKey) (txBytes []byte, payloadsToSign []*rosettatypes.SigningPayload, err error)
	// Tx converts a tendermint transaction and tx result if provided to a rosetta tx
	Tx(rawTx tmtypes.Tx, txResult *abci.ResponseDeliverTx) (*rosettatypes.Transaction, error)
	// TxIdentifiers converts a tendermint tx to transaction identifiers
	TxIdentifiers(txs []tmtypes.Tx) []*rosettatypes.TransactionIdentifier
	// BalanceOps converts events to balance operations
	BalanceOps(status string, events []abci.Event) []*rosettatypes.Operation
	// SyncStatus converts a tendermint status to sync status
	SyncStatus(status *tmcoretypes.ResultStatus) *rosettatypes.SyncStatus
	// Peers converts tendermint peers to rosetta
	Peers(peers []tmcoretypes.Peer) []*rosettatypes.Peer
}

// ToSDKConverter is an interface that exposes
// all the functions used to convert rosetta types
// to tendermint and sdk types
type ToSDKConverter interface {
	// UnsignedTx converts rosetta operations to an unsigned cosmos sdk transactions
	UnsignedTx(ops []*rosettatypes.Operation) (tx authsigning.Tx, err error)
	// SignedTx adds the provided signatures after decoding the unsigned transaction raw bytes
	// and returns the signed tx bytes
	SignedTx(txBytes []byte, signatures []*rosettatypes.Signature) (signedTxBytes []byte, err error)
	// Msg converts metadata to an sdk message
	Msg(meta map[string]interface{}, msg sdk.Msg) (err error)
	// HashToTxType returns the transaction type (end block, begin block or deliver tx)
	// and the real hash to query in order to get information
	HashToTxType(hashBytes []byte) (txType TransactionType, realHash []byte)
	// PubKey attempts to convert a rosetta public key to cosmos sdk one
	PubKey(pk *rosettatypes.PublicKey) (cryptotypes.PubKey, error)
}

type converter struct {
	newTxBuilder    func() sdkclient.TxBuilder
	txBuilderFromTx func(tx sdk.Tx) (sdkclient.TxBuilder, error)
	txDecode        sdk.TxDecoder
	txEncode        sdk.TxEncoder
	bytesToSign     func(tx authsigning.Tx, signerData authsigning.SignerData) (b []byte, err error)
	ir              codectypes.InterfaceRegistry
	cdc             *codec.ProtoCodec
}

func NewConverter(cdc *codec.ProtoCodec, ir codectypes.InterfaceRegistry, cfg sdkclient.TxConfig) Converter {
	_ = "STUB: not implemented"
	return *new(Converter)
}

func (c converter) ToSDK() ToSDKConverter { _ = "STUB: not implemented"; return *new(ToSDKConverter) }

func (c converter) ToRosetta() ToRosettaConverter {
	_ = "STUB: not implemented"

	// OpsToUnsignedTx returns all the sdk.Msgs given the operations
	return *new(ToRosettaConverter)
}

func (c converter) UnsignedTx(ops []*rosettatypes.Operation) (tx authsigning.Tx, err error) {
	_ = "STUB: not implemented"
	return *new(authsigning.Tx), nil
}

// verify message correctness

// check if there are enough signers

// append the msg

// if there's only one signer then simply continue

// after we have got the msg, we need to verify if the message has multiple signers
// if it has got multiple signers, then we need to fetch all the related operations
// which involve the other signers of the msg, we expect to find them in order
// so if the msg is named "v1.test.Send" and it expects 3 signers, the next 3 operations
// must be with the same name "v1.test.Send" and contain the other signers
// then we can just skip their processing

// get the next index
// verify that the operation is equal to the new one

// increase so we skip it

// Msg unmarshals the rosetta metadata to the given sdk.Msg
func (c converter) Msg(meta map[string]interface{}, msg sdk.Msg) error {
	_ = "STUB: not implemented"
	return nil
}

func (c converter) Meta(msg sdk.Msg) (meta map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ops will create an operation for each msg signer
// with the message proto name as type, and the raw fields
// as metadata
func (c converter) Ops(status string, msg sdk.Msg) ([]*rosettatypes.Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tx converts a tendermint raw transaction and its result (if provided) to a rosetta transaction
func (c converter) Tx(rawTx tmtypes.Tx, txResult *abci.ResponseDeliverTx) (*rosettatypes.Transaction, error) {
	_ = "STUB: not implemented"
	// decode tx
	return nil, nil
}

// get initial status, as per sdk design, if one msg fails
// the whole TX will be considered failing, so we can't have
// 1 msg being success and 1 msg being reverted

// if nil, we're probably checking an unconfirmed tx
// or trying to build a new transaction, so status
// is not put inside

// set the status

// get operations from msgs

// now get balance events from response deliver tx

// tx result might be nil, in case we're querying an unconfirmed tx from the mempool

// now normalize indexes

func (c converter) BalanceOps(status string, events []abci.Event) []*rosettatypes.Operation {
	_ = "STUB: not implemented"
	return nil
}

// sdkEventToBalanceOperations converts an event to a rosetta balance operation
// it will panic if the event is malformed because it might mean the sdk spec
// has changed and rosetta needs to reflect those changes too.
// The balance operations are multiple, one for each denom.
func sdkEventToBalanceOperations(status string, event abci.Event) (operations []*rosettatypes.Operation, isBalanceEvent bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// rosetta does not have the concept of burning coins, so we need to mock
// the burn as a send to an address that cannot be resolved to anything

// in case the event is a subtract balance one the rewrite value with
// the negative coin identifier

// Amounts converts []sdk.Coin to rosetta amounts
func (c converter) Amounts(ownedCoins []sdk.Coin, availableCoins sdk.Coins) []*rosettatypes.Amount {
	_ = "STUB: not implemented"
	return nil
}

// AddOperationIndexes adds the indexes to operations adhering to specific rules:
// operations related to messages will be always before than the balance ones
func AddOperationIndexes(msgOps []*rosettatypes.Operation, balanceOps []*rosettatypes.Operation) (finalOps []*rosettatypes.Operation) {
	_ = "STUB: not implemented"
	return nil
}

// add indexes to msg ops

// add indexes to balance ops

// EndBlockTxHash produces a mock endblock hash that rosetta can query
// for endblock operations, it also serves the purpose of representing
// part of the state changes happening at endblock level (balance ones)
func (c converter) EndBlockTxHash(hash []byte) string { _ = "STUB: not implemented"; return "" }

// BeginBlockTxHash produces a mock beginblock hash that rosetta can query
// for beginblock operations, it also serves the purpose of representing
// part of the state changes happening at beginblock level (balance ones)
func (c converter) BeginBlockTxHash(hash []byte) string { _ = "STUB: not implemented"; return "" }

// HashToTxType takes the provided hash bytes from rosetta and discerns if they are
// a deliver tx type or endblock/begin block hash, returning the real hash afterwards
func (c converter) HashToTxType(hashBytes []byte) (txType TransactionType, realHash []byte) {
	_ = "STUB: not implemented"
	return *new(TransactionType), nil
}

// StatusToSyncStatus converts a tendermint status to rosetta sync status
func (c converter) SyncStatus(status *tmcoretypes.ResultStatus) *rosettatypes.SyncStatus {
	_ = "STUB: not implemented"
	// determine sync status
	return nil
}

// sync info does not allow us to get target height

// TxIdentifiers converts a tendermint raw transactions into an array of rosetta tx identifiers
func (c converter) TxIdentifiers(txs []tmtypes.Tx) []*rosettatypes.TransactionIdentifier {
	_ = "STUB: not implemented"
	return nil
}

// tmResultBlockToRosettaBlockResponse converts a tendermint result block to block response
func (c converter) BlockResponse(block *tmcoretypes.ResultBlock) crgtypes.BlockResponse {
	_ = "STUB: not implemented"
	return *new(crgtypes.BlockResponse)
}

// Peers converts tm peers to rosetta peers
func (c converter) Peers(peers []tmcoretypes.Peer) []*rosettatypes.Peer {
	_ = "STUB: not implemented"
	return nil
}

// OpsAndSigners takes transactions bytes and returns the operation, is signed is true it will return
// the account identifiers which have signed the transaction
func (c converter) OpsAndSigners(txBytes []byte) (ops []*rosettatypes.Operation, signers []*rosettatypes.AccountIdentifier, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// get the signers

func (c converter) SignedTx(txBytes []byte, signatures []*rosettatypes.Signature) (signedTxBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//

// TODO(fdymylja): here we should check that the public key matches...

func (c converter) PubKey(pubKey *rosettatypes.PublicKey) (cryptotypes.PubKey, error) {
	_ = "STUB: not implemented"
	return *new(cryptotypes.PubKey), nil
}

// SigningComponents takes a sdk tx and construction metadata and returns signable components
func (c converter) SigningComponents(tx authsigning.Tx, metadata *ConstructionMetadata, rosPubKeys []*rosettatypes.PublicKey) (txBytes []byte, payloadsToSign []*rosettatypes.SigningPayload, err error) {
	_ = "STUB: not implemented"
	// verify metadata correctness
	return nil, nil, nil
}

// assert the signers data provided in options are the same as the expected signing accounts
// and that the number of rosetta provided public keys equals the one of the signers

// add transaction metadata

// build signatures

// pub key ordering matters, in a future release this check might be relaxed

// assert that the provided public keys are correctly ordered
// by checking if the signer at index i matches the pubkey at index

// set the signer data

// get signature bytes

// set payload

// set partial signature

// needs to be set to empty otherwise the codec will cry

// now we set the partial signatures in the tx
// because we will need to decode the sequence
// information of each account in a stateless way

// finally encode the tx

// SignerData converts the given any account to signer data
func (c converter) SignerData(anyAccount *codectypes.Any) (*SignerData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
