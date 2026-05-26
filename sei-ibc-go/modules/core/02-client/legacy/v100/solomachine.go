package v100

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/sei-protocol/sei-chain/sei-cosmos/codec"
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	sdk "github.com/sei-protocol/sei-chain/sei-cosmos/types"

	"github.com/sei-protocol/sei-chain/sei-ibc-go/modules/core/exported"
)

// NOTE: this is a mock implmentation for exported.ClientState. This implementation
// should only be registered on the InterfaceRegistry during cli command genesis migration.
// This implementation is only used to successfully unmarshal the previous solo machine
// client state and consensus state and migrate them to the new implementations. When the proto
// codec unmarshals, it calls UnpackInterfaces() to create a cached value of the any. The
// UnpackInterfaces function for IdenitifiedClientState will attempt to unpack the any to
// exported.ClientState. If the solomachine v1 type is not registered against the exported.ClientState
// the unmarshal will fail. This implementation will panic on every interface function.
// The same is done for the ConsensusState.

// Interface implementation checks.
var (
	_, _ codectypes.UnpackInterfacesMessage = &ClientState{}, &ConsensusState{}
	_    exported.ClientState               = (*ClientState)(nil)
	_    exported.ConsensusState            = &ConsensusState{}
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) { _ = "STUB: not implemented"; return }

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (cs ClientState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (cs ConsensusState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// ClientType panics!
func (cs ClientState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetLatestHeight panics!
func (cs ClientState) GetLatestHeight() exported.Height {
	_ = "STUB: not implemented"
	return *new(exported.Height)
}

// Status panics!
func (cs ClientState) Status(_ sdk.Context, _ sdk.KVStore, _ codec.BinaryCodec) exported.Status {
	_ = "STUB: not implemented"
	return *new(exported.Status)
}

// Validate panics!
func (cs ClientState) Validate() error { _ = "STUB: not implemented"; return nil }

// GetProofSpecs panics!
func (cs ClientState) GetProofSpecs() []*ics23.ProofSpec { _ = "STUB: not implemented"; return nil }

// ZeroCustomFields panics!
func (cs ClientState) ZeroCustomFields() exported.ClientState {
	_ = "STUB: not implemented"
	return *new(exported.ClientState)
}

// Initialize panics!
func (cs ClientState) Initialize(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, consState exported.ConsensusState) error {
	_ = "STUB: not implemented"
	return nil
}

// ExportMetadata panics!
func (cs ClientState) ExportMetadata(_ sdk.KVStore) []exported.GenesisMetadata {
	_ = "STUB: not implemented"
	return nil
}

// CheckHeaderAndUpdateState panics!
func (cs *ClientState) CheckHeaderAndUpdateState(
	_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Header,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// CheckMisbehaviourAndUpdateState panics!
func (cs ClientState) CheckMisbehaviourAndUpdateState(
	_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Misbehaviour,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// CheckSubstituteAndUpdateState panics!
func (cs ClientState) CheckSubstituteAndUpdateState(
	ctx sdk.Context, _ codec.BinaryCodec, _, _ sdk.KVStore,
	_ exported.ClientState,
) (exported.ClientState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), nil
}

// VerifyUpgradeAndUpdateState panics!
func (cs ClientState) VerifyUpgradeAndUpdateState(
	_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore,
	_ exported.ClientState, _ exported.ConsensusState, _, _ []byte,
) (exported.ClientState, exported.ConsensusState, error) {
	_ = "STUB: not implemented"
	return *new(exported.ClientState), *new(exported.ConsensusState), nil
}

// VerifyClientState panics!
func (cs ClientState) VerifyClientState(
	store sdk.KVStore, cdc codec.BinaryCodec,
	_ exported.Height, _ exported.Prefix, _ string, _ []byte, clientState exported.ClientState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyClientConsensusState panics!
func (cs ClientState) VerifyClientConsensusState(
	sdk.KVStore, codec.BinaryCodec,
	exported.Height, string, exported.Height, exported.Prefix,
	[]byte, exported.ConsensusState,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyConnectionState panics!
func (cs ClientState) VerifyConnectionState(
	sdk.KVStore, codec.BinaryCodec, exported.Height,
	exported.Prefix, []byte, string, exported.ConnectionI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyChannelState panics!
func (cs ClientState) VerifyChannelState(
	sdk.KVStore, codec.BinaryCodec, exported.Height, exported.Prefix,
	[]byte, string, string, exported.ChannelI,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketCommitment panics!
func (cs ClientState) VerifyPacketCommitment(
	sdk.Context, sdk.KVStore, codec.BinaryCodec, exported.Height,
	uint64, uint64, exported.Prefix, []byte,
	string, string, uint64, []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketAcknowledgement panics!
func (cs ClientState) VerifyPacketAcknowledgement(
	sdk.Context, sdk.KVStore, codec.BinaryCodec, exported.Height,
	uint64, uint64, exported.Prefix, []byte,
	string, string, uint64, []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyPacketReceiptAbsence panics!
func (cs ClientState) VerifyPacketReceiptAbsence(
	sdk.Context, sdk.KVStore, codec.BinaryCodec, exported.Height,
	uint64, uint64, exported.Prefix, []byte,
	string, string, uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyNextSequenceRecv panics!
func (cs ClientState) VerifyNextSequenceRecv(
	sdk.Context, sdk.KVStore, codec.BinaryCodec, exported.Height,
	uint64, uint64, exported.Prefix, []byte,
	string, string, uint64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ClientType panics!
func (ConsensusState) ClientType() string { _ = "STUB: not implemented"; return "" }

// GetTimestamp panics!
func (cs ConsensusState) GetTimestamp() uint64 { _ = "STUB: not implemented"; return 0 }

// GetRoot panics!
func (cs ConsensusState) GetRoot() exported.Root {
	_ = "STUB: not implemented"
	return *new(exported.Root)
}

// ValidateBasic panics!
func (cs ConsensusState) ValidateBasic() error { _ = "STUB: not implemented"; return nil }
