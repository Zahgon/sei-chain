package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

// Run is the main loop, but returns an error
func Run(args []string) error { _ = "STUB: not implemented"; return nil }

// if RestartAfterUpgrade, we launch after a successful upgrade (only condition LaunchProcess returns nil)
