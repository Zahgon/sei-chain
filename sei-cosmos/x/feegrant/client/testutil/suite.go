package testutil

import (
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/feegrant"
)

const (
	oneYear  = 365 * 24 * 60 * 60
	tenHours = 10 * 60 * 60
	oneHour  = 60 * 60
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg          network.Config
	network      *network.Network
	addedGranter sdk.AccAddress
	addedGrantee sdk.AccAddress
	addedGrant   feegrant.Grant
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	_ = "STUB: not implemented"
	return nil
}

func (s *IntegrationTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

// createGrant creates a new basic allowance fee grant from granter to grantee.
func (s *IntegrationTestSuite) createGrant(granter, grantee sdk.Address) {
	_ = "STUB: not implemented"
	return
}

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdGetFeeGrant() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdGetFeeGrantsByGrantee() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestCmdGetFeeGrantsByGranter() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCmdFeeGrant() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestNewCmdRevokeFeegrant() { _ = "STUB: not implemented"; return }

// Create new fee grant specifically to test amino.

func (s *IntegrationTestSuite) TestTxWithFeeGrant() { _ = "STUB: not implemented"; return }

// creating an account manually (This account won't be exist in state)

// granted fee allowance for an account which is not in state and creating
// any tx with it by using --fee-account shouldn't fail

func (s *IntegrationTestSuite) TestFilteredFeeAllowance() { _ = "STUB: not implemented"; return }

// get filtered fee allowance and check info

// exec filtered fee allowance

func getFormattedExpiration(duration int64) string { _ = "STUB: not implemented"; return "" }
