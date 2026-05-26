package signing

import (
	codectypes "github.com/sei-protocol/sei-chain/sei-cosmos/codec/types"
	cryptotypes "github.com/sei-protocol/sei-chain/sei-cosmos/crypto/types"
)

// SignatureV2 is a convenience type that is easier to use in application logic
// than the protobuf SignerInfo's and raw signature bytes. It goes beyond the
// first sdk.Signature types by supporting sign modes and explicitly nested
// multi-signatures. It is intended to be used for both building and verifying
// signatures.
type SignatureV2 struct {
	// PubKey is the public key to use for verifying the signature
	PubKey cryptotypes.PubKey

	// Data is the actual data of the signature which includes SignMode's and
	// the signatures themselves for either single or multi-signatures.
	Data SignatureData

	// Sequence is the sequence of this account. Only populated in
	// SIGN_MODE_DIRECT.
	Sequence uint64
}

// SignatureDataToProto converts a SignatureData to SignatureDescriptor_Data.
// SignatureDescriptor_Data is considered an encoding type whereas SignatureData is used for
// business logic.
func SignatureDataToProto(data SignatureData) *SignatureDescriptor_Data {
	_ = "STUB: not implemented"
	return nil
}

// SignatureDataFromProto converts a SignatureDescriptor_Data to SignatureData.
// SignatureDescriptor_Data is considered an encoding type whereas SignatureData is used for
// business logic.
func SignatureDataFromProto(descData *SignatureDescriptor_Data) SignatureData {
	_ = "STUB: not implemented"
	return *new(SignatureData)
}

var _, _ codectypes.UnpackInterfacesMessage = &SignatureDescriptors{}, &SignatureDescriptor{}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (sds *SignatureDescriptors) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpackInterfaces implements the UnpackInterfaceMessages.UnpackInterfaces method
func (sd *SignatureDescriptor) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	_ = "STUB: not implemented"
	return nil
}
