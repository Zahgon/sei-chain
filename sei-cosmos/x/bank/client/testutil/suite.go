package testutil

import (
	"github.com/stretchr/testify/suite"

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

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetBalancesCmd() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryTotalSupply() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestGetCmdQueryDenomsMetadata() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewSendTxCmdGenOnly() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewSendTxCmd() { _ = "STUB: not implemented"; return }

func NewCoin(denom string, amount sdk.Int) *sdk.Coin { _ = "STUB: not implemented"; return nil }
