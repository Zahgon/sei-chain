package testutil

import (
	seiapp "github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/sei-cosmos/testutil/network"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/stretchr/testify/suite"
)

func NewIntegrationTestSuite(cfg network.Config) *IntegrationTestSuite {
	_ = "STUB: not implemented"
	return nil
}

type IntegrationTestSuite struct {
	suite.Suite

	app     *seiapp.App
	cfg     network.Config
	network *network.Network
	ctx     sdk.Context
}

func (s *IntegrationTestSuite) SetupSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TearDownSuite() { _ = "STUB: not implemented"; return }

func (s *IntegrationTestSuite) TestModuleVersionsCLI() { _ = "STUB: not implemented"; return }

// avoid printing as yaml from CLI command

// setup expected response

// append new line to match behaviour of PrintProto

// get actual module versions list response from cli
