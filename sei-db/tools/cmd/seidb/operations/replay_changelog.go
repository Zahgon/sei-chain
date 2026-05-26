package operations

import (
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/sei-db/db_engine/types"
	"github.com/sei-protocol/sei-chain/sei-db/proto"
)

var ssStore types.StateStore

func ReplayChangelogCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func executeReplayChangelog(cmd *cobra.Command, _ []string) { _ = "STUB: not implemented"; return }

// use first available offset

// use latest offset

// open the database if this is not a dry run

// replay the changelog

// close the database

func processChangelogEntry(index uint64, entry proto.ChangelogEntry) error {
	_ = "STUB: not implemented"
	return nil
}
