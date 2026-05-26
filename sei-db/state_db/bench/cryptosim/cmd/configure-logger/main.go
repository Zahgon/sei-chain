// configure-logger reads a cryptosim config file and prints shell export
// statements that configure seilog's environment variables. Intended to be
// called via eval in a shell script before launching the benchmark binary.
//
// Usage:
//
//	eval "$(configure-logger config.json)"
package main

import (
	"fmt"
	"os"
)

/*
This extra binary for setting up logging is an unfortunate complexity, but we can't really avoid it
given that the only way to configure logging is to set environment variables before the main process starts.
*/

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "configure-logger: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// shellQuote wraps s in single quotes, escaping any embedded single quotes.
func shellQuote(s string) string { _ = "STUB: not implemented"; return "" }
