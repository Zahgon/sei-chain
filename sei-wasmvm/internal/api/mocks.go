package api

import (
	"encoding/json"
	"testing"

	"github.com/sei-protocol/sei-chain/sei-wasmvm/internal/api/testdb"
	"github.com/sei-protocol/sei-chain/sei-wasmvm/types"
)

/** helper constructors **/

const MOCK_CONTRACT_ADDR = "contract"

func MockEnv() types.Env { _ = "STUB: not implemented"; return *new(types.Env) }

func MockEnvBin(t *testing.T) []byte { _ = "STUB: not implemented"; return nil }

func MockInfo(sender types.HumanAddress, funds []types.Coin) types.MessageInfo {
	_ = "STUB: not implemented"
	return *new(types.MessageInfo)
}

func MockInfoWithFunds(sender types.HumanAddress) types.MessageInfo {
	_ = "STUB: not implemented"
	return *new(types.MessageInfo)
}

func MockInfoBin(t *testing.T, sender types.HumanAddress) []byte {
	_ = "STUB: not implemented"
	return nil
}

func MockIBCChannel(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannel {
	_ = "STUB: not implemented"
	return *new(types.IBCChannel)
}

func MockIBCChannelOpenInit(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannelOpenMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCChannelOpenMsg)
}

func MockIBCChannelOpenTry(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannelOpenMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCChannelOpenMsg)
}

func MockIBCChannelConnectAck(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannelConnectMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCChannelConnectMsg)
}

func MockIBCChannelConnectConfirm(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannelConnectMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCChannelConnectMsg)
}

func MockIBCChannelCloseInit(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannelCloseMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCChannelCloseMsg)
}

func MockIBCChannelCloseConfirm(channelID string, ordering types.IBCOrder, ibcVersion string) types.IBCChannelCloseMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCChannelCloseMsg)
}

func MockIBCPacket(myChannel string, data []byte) types.IBCPacket {
	_ = "STUB: not implemented"
	return *new(types.IBCPacket)
}

func MockIBCPacketReceive(myChannel string, data []byte) types.IBCPacketReceiveMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCPacketReceiveMsg)
}

func MockIBCPacketAck(myChannel string, data []byte, ack types.IBCAcknowledgement) types.IBCPacketAckMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCPacketAckMsg)
}

func MockIBCPacketTimeout(myChannel string, data []byte) types.IBCPacketTimeoutMsg {
	_ = "STUB: not implemented"
	return *new(types.IBCPacketTimeoutMsg)
}

/*** Mock GasMeter ****/
// This code is borrowed from Cosmos-SDK store/types/gas.go

// ErrorOutOfGas defines an error thrown when an action results in out of gas.
type ErrorOutOfGas struct {
	Descriptor string
}

// ErrorGasOverflow defines an error thrown when an action results gas consumption
// unsigned integer overflow.
type ErrorGasOverflow struct {
	Descriptor string
}

type MockGasMeter interface {
	types.GasMeter
	ConsumeGas(amount types.Gas, descriptor string)
}

type mockGasMeter struct {
	limit    types.Gas
	consumed types.Gas
}

// NewMockGasMeter returns a reference to a new mockGasMeter.
func NewMockGasMeter(limit types.Gas) MockGasMeter {
	_ = "STUB: not implemented"
	return *new(MockGasMeter)
}

func (g *mockGasMeter) GasConsumed() types.Gas { _ = "STUB: not implemented"; return *new(types.Gas) }

func (g *mockGasMeter) Limit() types.Gas {
	_ = "STUB: not implemented"

	// addUint64Overflow performs the addition operation on two uint64 integers and
	// returns a boolean on whether or not the result overflows.
	return *new(types.Gas)
}

func addUint64Overflow(a, b uint64) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (g *mockGasMeter) ConsumeGas(amount types.Gas, descriptor string) {
	_ = "STUB: not implemented"

	// TODO: Should we set the consumed field after overflow checking?
	return
}

/*** Mock types.KVStore ****/
// Much of this code is borrowed from Cosmos-SDK store/transient.go

// Note: these gas prices are all in *wasmer gas* and (sdk gas * 100)
//
// We making simple values and non-clear multiples so it is easy to see their impact in test output
// Also note we do not charge for each read on an iterator (out of simplicity and not needed for tests)
const (
	GetPrice    uint64 = 99000
	SetPrice    uint64 = 187000
	RemovePrice uint64 = 142000
	RangePrice  uint64 = 261000
)

type Lookup struct {
	db    *testdb.MemDB
	meter MockGasMeter
}

func NewLookup(meter MockGasMeter) *Lookup { _ = "STUB: not implemented"; return nil }

func (l *Lookup) SetGasMeter(meter MockGasMeter) { _ = "STUB: not implemented"; return }

func (l *Lookup) WithGasMeter(meter MockGasMeter) *Lookup { _ = "STUB: not implemented"; return nil }

// Get wraps the underlying DB's Get method panicing on error.
func (l Lookup) Get(key []byte) []byte { _ = "STUB: not implemented"; return nil }

// Set wraps the underlying DB's Set method panicing on error.
func (l Lookup) Set(key, value []byte) { _ = "STUB: not implemented"; return }

// Delete wraps the underlying DB's Delete method panicing on error.
func (l Lookup) Delete(key []byte) { _ = "STUB: not implemented"; return }

// Iterator wraps the underlying DB's Iterator method panicing on error.
func (l Lookup) Iterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

// ReverseIterator wraps the underlying DB's ReverseIterator method panicing on error.
func (l Lookup) ReverseIterator(start, end []byte) types.Iterator {
	_ = "STUB: not implemented"
	return *new(types.Iterator)
}

var _ types.KVStore = (*Lookup)(nil)

/***** Mock types.GoAPI ****/

const CanonicalLength = 32

const (
	CostCanonical uint64 = 440
	CostHuman     uint64 = 550
)

func MockCanonicalAddress(human string) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func MockHumanAddress(canon []byte) (string, uint64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func NewMockAPI() *types.GoAPI { _ = "STUB: not implemented"; return nil }

func TestMockApi(t *testing.T) { _ = "STUB: not implemented"; return }

/**** MockQuerier ****/

const DEFAULT_QUERIER_GAS_LIMIT = 1_000_000

type MockQuerier struct {
	Bank    BankQuerier
	Custom  CustomQuerier
	usedGas uint64
}

var _ types.Querier = &MockQuerier{}

func DefaultQuerier(contractAddr string, coins types.Coins) types.Querier {
	_ = "STUB: not implemented"
	return *new(types.Querier)
}

func (q *MockQuerier) Query(request types.QueryRequest, _gasLimit uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q MockQuerier) GasConsumed() uint64 { _ = "STUB: not implemented"; return 0 }

type BankQuerier struct {
	Balances map[string]types.Coins
}

func NewBankQuerier(balances map[string]types.Coins) BankQuerier {
	_ = "STUB: not implemented"
	return *new(BankQuerier)
}

func (q BankQuerier) Query(request *types.BankQuery) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CustomQuerier interface {
	Query(request json.RawMessage) ([]byte, error)
}

type NoCustom struct{}

var _ CustomQuerier = NoCustom{}

func (q NoCustom) Query(request json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReflectCustom fulfills the requirements for testing `reflect` contract
type ReflectCustom struct{}

var _ CustomQuerier = ReflectCustom{}

type CustomQuery struct {
	Ping        *struct{}         `json:"ping,omitempty"`
	Capitalized *CapitalizedQuery `json:"capitalized,omitempty"`
}

type CapitalizedQuery struct {
	Text string `json:"text"`
}

// CustomResponse is the response for all `CustomQuery`s
type CustomResponse struct {
	Msg string `json:"msg"`
}

func (q ReflectCustom) Query(request json.RawMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//************ test code for mocks *************************//

func TestBankQuerierAllBalances(t *testing.T) { _ = "STUB: not implemented"; return }

// query existing account

// query missing account

func TestBankQuerierBalance(t *testing.T) { _ = "STUB: not implemented"; return }

// query existing account with matching denom

// query existing account with missing denom

// query missing account

func TestReflectCustomQuerier(t *testing.T) { _ = "STUB: not implemented"; return }

// try ping

// try capital
