package types

import (
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// GenNodeID generates a random NodeID.
func GenNodeID(rng utils.Rng) NodeID { _ = "STUB: not implemented"; return *new(NodeID) }

// GenPublicKey generates a random PublicKey.
func GenPublicKey(rng utils.Rng) PublicKey { _ = "STUB: not implemented"; return *new(PublicKey) }

// GenSecretKey generates a random SecretKey.
func GenSecretKey(rng utils.Rng) SecretKey { _ = "STUB: not implemented"; return *new(SecretKey) }

// GenCommittee generates a random Committee of the given size.
// Returns the generated secret keys as well.
func GenCommittee(rng utils.Rng, size int) (*Committee, []SecretKey) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TestSecretKey creates a SecretKey for testing purposes.
// It uses NodeID as the seed of the secret key.
func TestSecretKey(nodeID NodeID) SecretKey { _ = "STUB: not implemented"; return *new(SecretKey) }

// GenLaneID generates a random LaneID.
func GenLaneID(rng utils.Rng) LaneID { _ = "STUB: not implemented"; return *new(LaneID) }

// GenSignature generates a random Signature.
func GenSignature(rng utils.Rng) *Signature { _ = "STUB: not implemented"; return nil }

// GenBlockNumber generates a random BlockNumber.
func GenBlockNumber(rng utils.Rng) BlockNumber { _ = "STUB: not implemented"; return *new(BlockNumber) }

// GenLaneRange generates a random LaneRange.
func GenLaneRange(rng utils.Rng) *LaneRange { _ = "STUB: not implemented"; return nil }

// GenBlockHeaderHash generates a random BlockHeaderHash.
func GenBlockHeaderHash(rng utils.Rng) BlockHeaderHash {
	_ = "STUB: not implemented"
	return *new(BlockHeaderHash)
}

// GenPayloadHash generates a random PayloadHash.
func GenPayloadHash(rng utils.Rng) PayloadHash { _ = "STUB: not implemented"; return *new(PayloadHash) }

// GenBlockHeader generates a random BlockHeader.
func GenBlockHeader(rng utils.Rng) *BlockHeader { _ = "STUB: not implemented"; return nil }

// GenPayload generates a random Payload.
func GenPayload(rng utils.Rng) *Payload { _ = "STUB: not implemented"; return nil }

// GenBlock generates a random Block.
func GenBlock(rng utils.Rng) *Block { _ = "STUB: not implemented"; return nil }

// GenSigned generates a random Signed.
func GenSigned[T Msg](rng utils.Rng, msg T) *Signed[T] { _ = "STUB: not implemented"; return nil }

// GenLaneProposal generates a random LaneProposal.
func GenLaneProposal(rng utils.Rng) *LaneProposal { _ = "STUB: not implemented"; return nil }

// GenLaneVote generates a random LaneVote.
func GenLaneVote(rng utils.Rng) *LaneVote { _ = "STUB: not implemented"; return nil }

// GenLaneQC generates a random LaneQC.
func GenLaneQC(rng utils.Rng) *LaneQC { _ = "STUB: not implemented"; return nil }

// GenRoadIndex generates a random RoadIndex.
func GenRoadIndex(rng utils.Rng) RoadIndex { _ = "STUB: not implemented"; return *new(RoadIndex) }

// GenViewNumber generates a random ViewNumber.
func GenViewNumber(rng utils.Rng) ViewNumber { _ = "STUB: not implemented"; return *new(ViewNumber) }

// GenView generates a random View.
func GenView(rng utils.Rng) View { _ = "STUB: not implemented"; return *new(View) }

// GenProposal generates a random Proposal.
func GenProposal(rng utils.Rng) *Proposal { _ = "STUB: not implemented"; return nil }

// GenProposalAt generates a Proposal at a specific view.
func GenProposalAt(rng utils.Rng, view View) *Proposal { _ = "STUB: not implemented"; return nil }

// GenAppHash generates a random AppHash.
func GenAppHash(rng utils.Rng) AppHash { _ = "STUB: not implemented"; return *new(AppHash) }

// GenAppProposal generates a random AppProposal.
func GenAppProposal(rng utils.Rng) *AppProposal { _ = "STUB: not implemented"; return nil }

// GenAppVote generates a random AppVote.
func GenAppVote(rng utils.Rng) *AppVote { _ = "STUB: not implemented"; return nil }

// GenAppQC generates a random AppQC.
func GenAppQC(rng utils.Rng) *AppQC { _ = "STUB: not implemented"; return nil }

// GenFullProposal generates a random FullProposal.
func GenFullProposal(rng utils.Rng) *FullProposal { _ = "STUB: not implemented"; return nil }

// GenGlobalBlockNumber generates a random GlobalBlockNumber.
func GenGlobalBlockNumber(rng utils.Rng) GlobalBlockNumber {
	_ = "STUB: not implemented"
	return *new(GlobalBlockNumber)
}

// GenGlobalBlock generates a random GlobalBlock.
func GenGlobalBlock(rng utils.Rng) *GlobalBlock { _ = "STUB: not implemented"; return nil }

// GenPrepareVote generates a random PrepareVote.
func GenPrepareVote(rng utils.Rng) *PrepareVote { _ = "STUB: not implemented"; return nil }

// GenPrepareQC generates a random PrepareQC.
func GenPrepareQC(rng utils.Rng) *PrepareQC { _ = "STUB: not implemented"; return nil }

// GenCommitVote generates a random CommitVote.
func GenCommitVote(rng utils.Rng) *CommitVote { _ = "STUB: not implemented"; return nil }

// GenCommitQC generates a random CommitQC.
func GenCommitQC(rng utils.Rng) *CommitQC { _ = "STUB: not implemented"; return nil }

// GenFullCommitQC generates a random FullCommitQC.
func GenFullCommitQC(rng utils.Rng) *FullCommitQC { _ = "STUB: not implemented"; return nil }

// GenTimeoutVote generates a random TimeoutVote.
func GenTimeoutVote(rng utils.Rng) *TimeoutVote { _ = "STUB: not implemented"; return nil }

// GenFullTimeoutVote generates a random FullTimeoutVote.
func GenFullTimeoutVote(rng utils.Rng) *FullTimeoutVote { _ = "STUB: not implemented"; return nil }

// GenTimeoutQC generates a random TimeoutQC.
func GenTimeoutQC(rng utils.Rng) *TimeoutQC { _ = "STUB: not implemented"; return nil }
