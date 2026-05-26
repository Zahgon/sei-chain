// Package nullify provides methods to init nil values structs for test assertion.
package nullify

import (
	"reflect"

	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"
)

var (
	coinType  = reflect.TypeOf(sdk.Coin{})
	coinsType = reflect.TypeOf(sdk.Coins{})
)

// Fill analyze all struct fields and slices with
// reflection and initialize the nil and empty slices,
// structs, and pointers.
func Fill(x interface{}) interface{} { _ = "STUB: not implemented"; return nil }

//nolint:gosec

//nolint:gosec
