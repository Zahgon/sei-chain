package debug

import (
	"github.com/spf13/cobra"
)

const (
	flagNodeRPCAddr = "rpc-laddr"
	flagProfAddr    = "pprof-laddr"
	flagFrequency   = "frequency"
)

func GetDebugCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }
