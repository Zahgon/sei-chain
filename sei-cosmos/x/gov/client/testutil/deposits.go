package testutil

import (
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
	"github.com/stretchr/testify/suite"
)

type DepositTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	fees    string
}

func NewDepositTestSuite(cfg network.Config) *DepositTestSuite {
	_ = "STUB: not implemented"
	return nil
}

func (s *DepositTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

func (s *DepositTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *DepositTestSuite) TestQueryDepositsInitialDeposit() { _ = "STUB: not implemented"; return }

// create a proposal with deposit

// deposit more amount

// waiting for voting period to end

// query deposit & verify initial deposit

// query deposits

// verify initial deposit

func (s *DepositTestSuite) TestQueryDepositsWithoutInitialDeposit() {
	_ = "STUB: not implemented"
	return
}

// create a proposal without deposit

// deposit amount

// waiting for voting period to end

// query deposit

// query deposits

// verify initial deposit

func (s *DepositTestSuite) TestQueryProposalNotEnoughDeposits() { _ = "STUB: not implemented"; return }

// create a proposal with deposit

// query proposal

// waiting for deposit period to end

// query proposal

func (s *DepositTestSuite) TestRejectedProposalDeposits() { _ = "STUB: not implemented"; return }

// create a proposal with deposit

// query deposits

// verify initial deposit

// vote

// query deposits

// verify initial deposit

func (s *DepositTestSuite) queryDeposits(val *network.Validator, proposalID string, exceptErr bool) types.Deposits {
	_ = "STUB: not implemented"
	return *new(types.Deposits)
}

func (s *DepositTestSuite) queryDeposit(val *network.Validator, proposalID string, exceptErr bool) *types.Deposit {
	_ = "STUB: not implemented"
	return nil
}
