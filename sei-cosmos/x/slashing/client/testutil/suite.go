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

// SetupSuite executes bootstrapping logic before all the tests, i.e. once before
// the entire suite, start executing.
func (s *IntegrationTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

// TearDownSuite performs cleanup logic after all the tests, i.e. once after the
// entire suite, has finished executing.
func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQuerySigningInfo() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryParams() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewUnjailTxCmd() { _ = "STUB: not implemented"; return }

// sync mode as there are no funds yet
