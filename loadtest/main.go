package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/std"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
	"github.com/sei-protocol/sei-chain/sei-cosmos/x/auth/tx"

	"github.com/sei-protocol/sei-chain/app"
)

var TestConfig EncodingConfig

const (
	VortexData = "{\"position_effect\":\"Open\",\"leverage\":\"1\"}"
)

var (
	FromMili                  = sdk.NewDec(1000000)
	producedCountPerMsgType   = make(map[string]*int64)
	sentCountPerMsgType       = make(map[string]*int64)
	prevSentCounterPerMsgType = make(map[string]*int64)

	BlockHeightsWithTxs = []int{}
	EvmTxHashes         = []common.Hash{}
)

type BlockData struct {
	Txs []string `json:"txs"`
}

type BlockHeader struct {
	Time string `json:"time"`
}

func init() {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	TestConfig = EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}
	std.RegisterLegacyAminoCodec(TestConfig.Amino)
	std.RegisterInterfaces(TestConfig.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(TestConfig.Amino)
	app.ModuleBasics.RegisterInterfaces(TestConfig.InterfaceRegistry)
	// Add this so that we don't end up getting disconnected for EVM client
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100
}

// deployEvmContract executes a bash script and returns its output as a string.
//
//nolint:gosec
func deployEvmContract(scriptPath string, config *Config) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

//nolint:gosec
func deployEvmContracts(config *Config) { _ = "STUB: not implemented"; return }

//TODO: this really should use `deployEvmContract`

func run(config *Config) { _ = "STUB: not implemented"; return }

// Start metrics collector in another thread

// initialize clients with addresses on clients

// starts loadtest workers. If config.Constant is true, then we don't gather loadtest results and let producer/consumer
// workers continue running. If config.Constant is false, then we will gather load test results in a file
func startLoadtestWorkers(client *LoadTestClient, config Config) { _ = "STUB: not implemented"; return }

// Catch OS signals for graceful shutdown

// Create producers and consumers

//nolint:gosec
//nolint:gosec

// Statistics reporting goroutine

//nolint:gosec

// Wait for a termination signal

func printStats(
	startTime time.Time,
	producedCountPerMsgType map[string]*int64,
	sentCountPerMsgType map[string]*int64,
	prevSentPerCounterPerMsgType map[string]*int64,
	blockHeights []int,
	blockTimes []string,
) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

//nolint:gosec

//nolint:gosec

// Generate a random message, only generate one admin message per block to prevent acc seq errors
func (c *LoadTestClient) getRandomMessageType(messageTypes []string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *LoadTestClient) generateMessage(key cryptotypes.PrivKey, msgType string) ([]sdk.Msg, bool, cryptotypes.PrivKey, uint64, int64) {
	_ = "STUB: not implemented"
	return nil, false, *new(cryptotypes.PrivKey), 0, 0
}

// TODO: Potentially just hard code the Funds amount here

// maybe make this configurable as well in the future

//nolint:gosec

// Randomly pick someone to redelegate / unbond from

// No denoms, let's mint

// generate some values for indices 1-100

//nolint:gosec

// generate random value

// nolint:gosec

// nolint:gosec

func (c *LoadTestClient) generateStakingMsg(delegatorAddr string, chosenValidator string, srcAddr string) sdk.Msg {
	_ = "STUB: not implemented"
	return *new(sdk.Msg)
}

// Randomly unbond, redelegate or delegate
// However, if there are no delegations, do so first

// Update delegation map

// nolint
func getLastHeight(blockchainEndpoint string) int { _ = "STUB: not implemented"; return 0 }

func getTxBlockInfo(blockchainEndpoint string, height string) (int, string, error) {
	_ = "STUB: not implemented"
	return 0, "", nil
}

func runEvmQueries(config Config) { _ = "STUB: not implemented"; return }

func GetDefaultConfigFilePath() string { _ = "STUB: not implemented"; return "" }

func ReadConfig(path string) Config { _ = "STUB: not implemented"; return *new(Config) }

func main() {
	configFilePath := flag.String("config-file", GetDefaultConfigFilePath(), "Path to the config.json file to use for this run")
	flag.Parse()

	config := ReadConfig(*configFilePath)
	fmt.Printf("Using config file: %s\n", *configFilePath)
	run(&config)
}
