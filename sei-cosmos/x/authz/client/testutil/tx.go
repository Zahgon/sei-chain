package testutil

import (
	"github.com/stretchr/testify/suite"

	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	bank "github.com/sei-protocol/sei-chain/sei-cosmos/x/bank/types"
	govtypes "github.com/sei-protocol/sei-chain/sei-cosmos/x/gov/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
	grantee sdk.AccAddress
}

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	_ = "STUB: not implemented"
	return nil
}

func (s *IntegrationTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

// Create new account in the keyring.

// Send some funds to the new account.

// create a proposal with deposit

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

var typeMsgSend = bank.SendAuthorization{}.MsgTypeURL()
var typeMsgVote = sdk.MsgTypeURL(&govtypes.MsgVote{})
var typeMsgSubmitProposal = sdk.MsgTypeURL(&govtypes.MsgSubmitProposal{})

func (s *IntegrationTestSuite) TestCLITxGrantAuthorization() { _ = "STUB: not implemented"; return }

// TODO: enable in v0.45

func execDelegate(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	_ = "STUB: not implemented"
	return *new(testutil.BufferWriter), nil
}

func (s *IntegrationTestSuite) TestCmdRevokeAuthorizations() { _ = "STUB: not implemented"; return }

// send-authorization

// generic-authorization

// generic-authorization used for amino testing

func (s *IntegrationTestSuite) TestExecAuthorizationWithExpiration() {
	_ = "STUB: not implemented"
	return
}

// msg vote

// waiting for authorization to expires

func (s *IntegrationTestSuite) TestNewExecGenericAuthorized() { _ = "STUB: not implemented"; return }

// msg vote

func (s *IntegrationTestSuite) TestNewExecGrantAuthorized() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestExecDelegateAuthorization() { _ = "STUB: not implemented"; return }

// test delegate no spend-limit

// test delegating to denied validator

func (s *IntegrationTestSuite) TestExecUndelegateAuthorization() { _ = "STUB: not implemented"; return }

// granting undelegate msg authorization

// delegating stakes to validator

// grant undelegate authorization without limit
