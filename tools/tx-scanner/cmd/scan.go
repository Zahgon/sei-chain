package cmd

import (
	"github.com/spf13/cobra"
)

func ScanCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func execute(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// Handle ALL the queries in a batch concurrently

// Wait for ALL queries in this batch to finish and then check any failures

// update the state

// processBlock processes a single block to find missing transactions
func processBlock(height int64) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Query the block to get the number of TXs

// Get all indexed TXs events

// Check if the number matches

// Now make sure each TX does exist

func getLatestBlockHeight() int64 { _ = "STUB: not implemented"; return 0 }
