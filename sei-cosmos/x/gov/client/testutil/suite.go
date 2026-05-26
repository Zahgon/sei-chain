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

// create a proposal with deposit

// vote for proposal

// create a proposal without deposit

// create a proposal3 with deposit

// vote for proposal3 as val

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdParams() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdParam() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdProposer() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdTally() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCmdSubmitProposal() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdGetProposal() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdGetProposals() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdQueryDeposits() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdQueryDeposit() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCmdDeposit() { _ = "STUB: not implemented"; return }

// 10stake

// 10stake

// 10stake

func (s *IntegrationTestSuite) TestCmdQueryVotes() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdQueryVote() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCmdVote() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCmdWeightedVote() { _ = "STUB: not implemented"; return }
