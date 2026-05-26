package cmd

import (
	_ "net/http/pprof" //nolint:gosec

	"github.com/spf13/cobra"

	ethtests "github.com/ethereum/go-ethereum/tests"
)

func BlocktestCmd(defaultNodeHome string) *cobra.Command { _ = "STUB: not implemented"; return nil }

// turn on Cancun for block test

func testIngester(testFilePath string, testName string) *ethtests.BlockTest {
	_ = "STUB: not implemented"
	return nil
}
