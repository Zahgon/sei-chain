package privval

import (
	"github.com/gogo/protobuf/proto"

	privvalproto "github.com/sei-protocol/sei-chain/sei-tendermint/proto/tendermint/privval"
)

// TODO: Add ChainIDRequest

func mustWrapMsg(pb proto.Message) privvalproto.Message {
	_ = "STUB: not implemented"
	return *new(privvalproto.Message)
}
