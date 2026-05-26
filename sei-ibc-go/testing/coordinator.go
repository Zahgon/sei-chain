package ibctesting

import (
	"testing"
	"time"
)

var (
	ChainIDPrefix   = "testchain"
	globalStartTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	TimeIncrement   = time.Second * 5
)

// Coordinator is a testing struct which contains N TestChain's. It handles keeping all chains
// in sync with regards to time.
type Coordinator struct {
	*testing.T

	CurrentTime time.Time
	Chains      map[string]*TestChain
}

// NewCoordinator initializes Coordinator with N TestChain's
func NewCoordinator(t *testing.T, n int) *Coordinator { _ = "STUB: not implemented"; return nil }

// IncrementTime iterates through all the TestChain's and increments their current header time
// by 5 seconds.
//
// CONTRACT: this function must be called after every Commit on any TestChain.
func (coord *Coordinator) IncrementTime() { _ = "STUB: not implemented"; return }

// IncrementTimeBy iterates through all the TestChain's and increments their current header time
// by specified time.
func (coord *Coordinator) IncrementTimeBy(increment time.Duration) {
	_ = "STUB: not implemented"
	return
}

// UpdateTime updates all clocks for the TestChains to the current global time.
func (coord *Coordinator) UpdateTime() { _ = "STUB: not implemented"; return }

// UpdateTimeForChain updates the clock for a specific chain.
func (coord *Coordinator) UpdateTimeForChain(chain *TestChain) { _ = "STUB: not implemented"; return }

// Setup constructs a TM client, connection, and channel on both chains provided. It will
// fail if any error occurs. The clientID's, TestConnections, and TestChannels are returned
// for both chains. The channels created are connected to the ibc-transfer application.
func (coord *Coordinator) Setup(path *Path) { _ = "STUB: not implemented"; return }

// channels can also be referenced through the returned connections

// SetupClients is a helper function to create clients on both chains. It assumes the
// caller does not anticipate any errors.
func (coord *Coordinator) SetupClients(path *Path) { _ = "STUB: not implemented"; return }

// SetupClientConnections is a helper function to create clients and the appropriate
// connections on both the source and counterparty chain. It assumes the caller does not
// anticipate any errors.
func (coord *Coordinator) SetupConnections(path *Path) { _ = "STUB: not implemented"; return }

// CreateConnection constructs and executes connection handshake messages in order to create
// OPEN channels on chainA and chainB. The connection information of for chainA and chainB
// are returned within a TestConnection struct. The function expects the connections to be
// successfully opened otherwise testing will fail.
func (coord *Coordinator) CreateConnections(path *Path) { _ = "STUB: not implemented"; return }

// ensure counterparty is up to date

// CreateMockChannels constructs and executes channel handshake messages to create OPEN
// channels that use a mock application module that returns nil on all callbacks. This
// function is expects the channels to be successfully opened otherwise testing will
// fail.
func (coord *Coordinator) CreateMockChannels(path *Path) { _ = "STUB: not implemented"; return }

// CreateTransferChannels constructs and executes channel handshake messages to create OPEN
// ibc-transfer channels on chainA and chainB. The function expects the channels to be
// successfully opened otherwise testing will fail.
func (coord *Coordinator) CreateTransferChannels(path *Path) { _ = "STUB: not implemented"; return }

// CreateChannel constructs and executes channel handshake messages in order to create
// OPEN channels on chainA and chainB. The function expects the channels to be successfully
// opened otherwise testing will fail.
func (coord *Coordinator) CreateChannels(path *Path) { _ = "STUB: not implemented"; return }

// ensure counterparty is up to date

// GetChain returns the TestChain using the given chainID and returns an error if it does
// not exist.
func (coord *Coordinator) GetChain(chainID string) *TestChain {
	_ = "STUB: not implemented"
	return nil
}

// GetChainID returns the chainID used for the provided index.
func GetChainID(index int) string { _ = "STUB: not implemented"; return "" }

// CommitBlock commits a block on the provided indexes and then increments the global time.
//
// CONTRACT: the passed in list of indexes must not contain duplicates
func (coord *Coordinator) CommitBlock(chains ...*TestChain) { _ = "STUB: not implemented"; return }

// CommitNBlocks commits n blocks to state and updates the block height by 1 for each commit.
func (coord *Coordinator) CommitNBlocks(chain *TestChain, n uint64) {
	_ = "STUB: not implemented"
	return
}

// ConnOpenInitOnBothChains initializes a connection on both endpoints with the state INIT
// using the OpenInit handshake call.
func (coord *Coordinator) ConnOpenInitOnBothChains(path *Path) error {
	_ = "STUB: not implemented"
	return nil
}

// ChanOpenInitOnBothChains initializes a channel on the source chain and counterparty chain
// with the state INIT using the OpenInit handshake call.
func (coord *Coordinator) ChanOpenInitOnBothChains(path *Path) error {
	_ = "STUB: not implemented"
	// NOTE: only creation of a capability for a transfer or mock port is supported
	// Other applications must bind to the port in InitGenesis or modify this code.
	return nil
}
