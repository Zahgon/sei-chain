package kv

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/pubsub/query/syntax"
)

func intInSlice(a int, list []int) bool { _ = "STUB: not implemented"; return false }

func int64FromBytes(bz []byte) int64 { _ = "STUB: not implemented"; return 0 }

func int64ToBytes(i int64) []byte { _ = "STUB: not implemented"; return nil }

func heightKey(height int64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func eventKey(compositeKey, typ, eventValue string, height int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseValueFromPrimaryKey(key []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func parseValueFromEventKey(key []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

func lookForHeight(conditions []syntax.Condition) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}
