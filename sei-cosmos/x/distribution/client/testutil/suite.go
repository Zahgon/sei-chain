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

// SetupTest creates a new network for _each_ integration test. We create a new
// network for each test because there are some state modifications that are
// needed to be made in order to make useful queries. However, we don't want
// these state changes to be present in other tests.
func (s *IntegrationTestSuite) SetupTest() { _ = "STUB: not implemented"; return }

// TearDownTest cleans up the curret test network after _each_ test.
func (s *IntegrationTestSuite) TearDownTest() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryParams() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryValidatorOutstandingRewards() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryValidatorCommission() {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TestGetCmdQueryValidatorSlashes() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryDelegatorRewards() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryCommunityPool() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewWithdrawRewardsCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewWithdrawAllRewardsCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewSetWithdrawAddrCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewFundCommunityPoolCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdSubmitProposal() { _ = "STUB: not implemented"; return }

// sync mode as there are no funds yet
