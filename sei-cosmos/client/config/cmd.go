package config

import (
	"github.com/spf13/cobra"
)

// Cmd returns a CLI command to interactively create an application CLI
// config file.
func Cmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runConfigCmd(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// print all client config fields to stdout

// it's a get

// it's set

func errUnknownConfigKey(key string) error { _ = "STUB: not implemented"; return nil }
