package main

import (
	"log/slog"

	"github.com/urfave/cli/v2"
)

// TODO (cody.littley): convert all commands to use flags stored in these variables
var (
	srcFlag = &cli.StringSliceFlag{
		Name:     "src",
		Aliases:  []string{"s"},
		Usage:    "Source paths where the DB data is found, at least one is required.",
		Required: true,
	}
	forceFlag = &cli.BoolFlag{
		Name:    "force",
		Aliases: []string{"f"},
		Usage:   "Force the operation without prompting for confirmation.",
	}
	knownHostsFileFlag = &cli.StringFlag{
		Name:     "known-hosts",
		Aliases:  []string{"k"},
		Usage:    "Path to a file containing known hosts for SSH connections.",
		Required: false,
		Value:    "~/.ssh/known_hosts",
	}
)

// buildCliParser creates a command line parser for the LittDB CLI tool.
func buildCLIParser(logger *slog.Logger) *cli.App { _ = "STUB: not implemented"; return nil }

// TODO (cody.littley) test in preprod

// Default to 0, meaning no age limit

// Builds a function that is called before any command is executed.
func buildBeforeAction(logger *slog.Logger) func(*cli.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If debug mode is enabled, this function will block until the user presses Enter.
func handleDebugMode(ctx *cli.Context, logger *slog.Logger) { _ = "STUB: not implemented"; return }

// block until newline is read

// If pprof is enabled, this function starts the pprof server.
func handlePProfMode(ctx *cli.Context, logger *slog.Logger) error {
	_ = "STUB: not implemented"
	return nil
}
