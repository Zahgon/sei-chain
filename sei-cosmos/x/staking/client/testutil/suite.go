package testutil

import (
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
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

// redelegate

// unbonding

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCreateValidatorCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryValidator() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryValidators() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryDelegation() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryDelegations() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryValidatorDelegations() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryUnbondingDelegations() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryUnbondingDelegation() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryValidatorUnbondingDelegations() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryRedelegations() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryRedelegation() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryValidatorRedelegations() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryHistoricalInfo() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryParams() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryPool() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewEditValidatorCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewDelegateCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewRedelegateCmd() { _ = "STUB: not implemented"; return }

// src-validator-addr
// dst-validator-addr

// src-validator-addr
// dst-validator-addr
// amount

// dst-validator-addr
// src-validator-addr
// amount

// src-validator-addr
// dst-validator-addr
// amount

func (s *IntegrationTestSuite) TestNewUnbondCmd() { _ = "STUB: not implemented"; return }

// TestBlockResults tests that the validator updates correctly show when
// calling the /block_results RPC endpoint.
// ref: https://github.com/cosmos/cosmos-sdk/issues/7401.
func (s *IntegrationTestSuite) TestBlockResults() { _ = "STUB: not implemented"; return }

// Create new account in the keyring.

// Send some funds to the new account.

// Use CLI to create a delegation from the new account to validator `val`.

// Create a HTTP rpc client.

// Loop until we find a block result with the correct validator updates.
// By experience, it happens around 2 blocks after `delHeight`.

// Wait maximum 10 blocks, or else fail test.

// We got our validator update, test passed.

// https://github.com/cosmos/cosmos-sdk/issues/10660
func (s *IntegrationTestSuite) TestEditValidatorMoniker() { _ = "STUB: not implemented"; return }
