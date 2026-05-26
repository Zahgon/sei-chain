package types

import (
	"fmt"
	"sort"
	"time"

	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/autobahn/pb"
	"github.com/sei-protocol/sei-chain/sei-tendermint/internal/protoutils"
	"github.com/sei-protocol/sei-chain/sei-tendermint/libs/utils"
)

// LaneRange represents a range [first,next) of blocks of a lane.
type LaneRange struct {
	utils.ReadOnly
	lane     LaneID
	first    BlockNumber
	next     BlockNumber
	lastHash BlockHeaderHash
}

// NewLaneRange constructs a LaneRange.
func NewLaneRange(lane LaneID, first BlockNumber, h utils.Option[*BlockHeader]) *LaneRange {
	_ = "STUB: not implemented"
	return nil
}

// Lane of this block range.
func (m *LaneRange) Lane() LaneID {
	_ = "STUB: not implemented"

	// First block of the range.
	return *new(LaneID)
}

func (m *LaneRange) First() BlockNumber {
	_ = "STUB: not implemented"

	// Next is the block after the last block of the range.
	return *new(BlockNumber)
}

func (m *LaneRange) Next() BlockNumber {
	_ = "STUB: not implemented"

	// Len returns the number of blocks in the range.
	return *new(BlockNumber)
}

func (m *LaneRange) Len() uint64 { _ = "STUB: not implemented"; return 0 }

// LastHash is the hash of the last block of the range.
// Returns a zero hash for an empty range.
func (m *LaneRange) LastHash() BlockHeaderHash {
	_ = "STUB: not implemented"

	// Verify verifies the LaneRange against the committee.
	return *new(BlockHeaderHash)
}

func (m *LaneRange) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// GlobalRange represents a [First,Next) range of global blocks.
type GlobalRange struct {
	First GlobalBlockNumber
	Next  GlobalBlockNumber
}

// Len returns the number of global blocks in the range.
func (g GlobalRange) Len() uint64 { _ = "STUB: not implemented"; return 0 }

func (g GlobalRange) Has(n GlobalBlockNumber) bool { _ = "STUB: not implemented"; return false }

// RoadIndex is the index of the consensus instance.
type RoadIndex uint64

// ViewNumber is the view number of the consensus instance.
type ViewNumber uint64

// Next view number.
func (n ViewNumber) Next() ViewNumber {
	_ = "STUB: not implemented"

	// View represents a consensus view.
	return *new(ViewNumber)
}

type View struct {
	Index  RoadIndex
	Number ViewNumber
}

// Less checks if v is earlier than b.
func (v View) Less(b View) bool { _ = "STUB: not implemented"; return false }

// Next returns the next view.
func (v View) Next() View { _ = "STUB: not implemented"; return *new(View) }

// ViewSpec is a justification to start a given view.
type ViewSpec struct {
	// WARNING: currently we have implicit assumption that
	// TimeoutQC.View().Index == CommitQC.Index.Next(),
	// I.e. that TimeoutQC comes from the expected consensus instance.
	CommitQC  utils.Option[*CommitQC]
	TimeoutQC utils.Option[*TimeoutQC]
}

// View is the view justified by vs.
func (vs *ViewSpec) View() View { _ = "STUB: not implemented"; return *new(View) }

func (vs *ViewSpec) NextTimestamp(c *Committee) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Proposal is the road tipcut proposal.
// It consists of ranges of blocks of each lane.
// AppQC could be nil if we haven't reached any quorum state hash.
type Proposal struct {
	utils.ReadOnly
	view       View
	timestamp  time.Time
	laneRanges map[LaneID]*LaneRange
	app        utils.Option[*AppProposal]
	// derived
	// WARNING: this is not a valid global range, because
	// it does not take into consideration committee.FirstBlock().
	// We keep it precomputed just to optimize the GlobalRange call.
	globalRangeWithoutOffset GlobalRange
}

func newProposal(view View, timestamp time.Time, laneRanges []*LaneRange, app utils.Option[*AppProposal]) *Proposal {
	_ = "STUB: not implemented"
	return nil
}

// Index of the proposal.
func (m *Proposal) Index() RoadIndex {
	_ = "STUB: not implemented"

	// View of the proposal.
	return *new(RoadIndex)
}

func (m *Proposal) View() View {
	_ = "STUB: not implemented"

	// Timestamp of the proposal.
	return *new(View)
}

func (m *Proposal) Timestamp() time.Time {
	_ = "STUB: not implemented"

	// App .
	return *new(time.Time)
}

func (m *Proposal) App() utils.Option[*AppProposal] {
	_ = "STUB: not implemented"

	// GlobalRange returns the proposed global block range.
	// To compute GlobalRange from lane ranges in proposal,
	// we need to know the global number of the first block
	// of the chain (c.FirstBlock()).
	return nil
}

func (m *Proposal) GlobalRange(c *Committee) GlobalRange {
	_ = "STUB: not implemented"
	return *new(GlobalRange)
}

// Arbitrary deterministic minimal diff between consecutive blocks.
const minTimestampDiff = time.Microsecond

// Monotone timestamp assigned to each block of the proposal.
// Returns None, if n doed not belong to the proposal's global range.
func (m *Proposal) BlockTimestamp(c *Committee, n GlobalBlockNumber) utils.Option[time.Time] {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec // TODO: do stricter timestamp validation before running in prod.

// Lowest allowed timestamp for the next index proposal.
func (m *Proposal) NextTimestamp() time.Time {
	_ = "STUB: not implemented"
	//nolint:gosec // TODO: do stricter timestamp validation before running in prod.
	return *new(time.Time)
}

// Verify checks that every present lane range belongs to the committee
// and is internally valid. Lanes may be omitted — omitted lanes are
// treated as implicit empty ranges by FullProposal.Verify.
func (m *Proposal) Verify(c *Committee) error { _ = "STUB: not implemented"; return nil }

// LaneRange returns the range of blocks of the given lane.
func (m *Proposal) LaneRange(lane LaneID) *LaneRange { _ = "STUB: not implemented"; return nil }

// FullProposal is a proposal with justification.
type FullProposal struct {
	utils.ReadOnly
	proposal  *Signed[*Proposal]
	laneQCs   map[LaneID]*LaneQC
	appQC     utils.Option[*AppQC]
	timeoutQC utils.Option[*TimeoutQC]
}

// NewReproposal creates a new reproposal based on viewSpec.
// Returns false if reproposal is not expected.
func NewReproposal(
	key SecretKey,
	viewSpec ViewSpec,
) (*FullProposal, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// NewProposal creates a new FullProposal.
// timestamp might get replaced to ensure that timestamps are monotone.
func NewProposal(
	key SecretKey,
	committee *Committee,
	viewSpec ViewSpec,
	timestamp time.Time,
	laneQCs map[LaneID]*LaneQC,
	appQC utils.Option[*AppQC],
) (*FullProposal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the new appProposal is not later than the previous one, then clear appQC.

// If the new appProposal is from the future (which may happen if this node is behind), then clear appQC.
// The proposal will be useless in this case, but at least it will be valid.

// Normalize the creation timestamp.

// Proposal .
func (m *FullProposal) Proposal() *Signed[*Proposal] {
	_ = "STUB: not implemented"

	// View .
	return nil
}

func (m *FullProposal) View() View { _ = "STUB: not implemented"; return *new(View) }

// LaneQC .
func (m *FullProposal) LaneQC(lane LaneID) (*LaneQC, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TimeoutQC returns the timeout QC if it exists.
func (m *FullProposal) TimeoutQC() utils.Option[*TimeoutQC] {
	_ = "STUB: not implemented"

	// Verify verifies the FullProposal against the current view.
	return nil
}

func (m *FullProposal) Verify(c *Committee, vs ViewSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Does the view match?

// Is the timestamp monotone?

// Is proposer valid?

// Verify the proposer's signature.

// Do we have the required timeoutQC?

// Verify timeoutQC.

// Is this a reproposal?

// Valid reproposal, no further verification needed.

// Verify the proposal's lane structure against the committee.

// Verify each lane range against the previous commitQC and its laneQC justification.

// Verify that range matches previous commitQC.

// Verify that the necessary laneQC is present and valid.

// Verify the appQC.

// LaneRangeConv is the protobuf converter for LaneRange.
var LaneRangeConv = protoutils.Conv[*LaneRange, *pb.LaneRange]{
	Encode: func(m *LaneRange) *pb.LaneRange {
		return &pb.LaneRange{
			Lane:     PublicKeyConv.Encode(m.lane),
			First:    utils.Alloc(uint64(m.first)),
			Next:     utils.Alloc(uint64(m.next)),
			LastHash: m.lastHash[:],
		}
	},
	Decode: func(m *pb.LaneRange) (*LaneRange, error) {
		lane, err := PublicKeyConv.DecodeReq(m.Lane)
		if err != nil {
			return nil, fmt.Errorf("Lane: %w", err)
		}
		if m.First == nil {
			return nil, fmt.Errorf("First: missing")
		}
		if m.Next == nil {
			return nil, fmt.Errorf("Next: missing")
		}
		lastHash, err := ParseBlockHeaderHash(m.LastHash)
		if err != nil {
			return nil, fmt.Errorf("LastHash: %w", err)
		}
		return &LaneRange{
			lane:     lane,
			first:    BlockNumber(*m.First),
			next:     BlockNumber(*m.Next),
			lastHash: lastHash,
		}, nil
	},
}

// ViewConv is the protobuf converter for View.
var ViewConv = protoutils.Conv[View, *pb.View]{
	Encode: func(m View) *pb.View {
		return &pb.View{
			Index:  utils.Alloc(uint64(m.Index)),
			Number: utils.Alloc(uint64(m.Number)),
		}
	},
	Decode: func(m *pb.View) (View, error) {
		if m.Index == nil {
			return View{}, fmt.Errorf("index: missing")
		}
		if m.Number == nil {
			return View{}, fmt.Errorf("number: missing")
		}
		return View{
			Index:  RoadIndex(*m.Index),
			Number: ViewNumber(*m.Number),
		}, nil
	},
}

// ProposalConv is the protobuf converter for Proposal.
var ProposalConv = protoutils.Conv[*Proposal, *pb.Proposal]{
	Encode: func(m *Proposal) *pb.Proposal {
		laneRanges := make([]*LaneRange, 0, len(m.laneRanges))
		for _, r := range m.laneRanges {
			laneRanges = append(laneRanges, r)
		}
		sort.Slice(laneRanges, func(i, j int) bool { return laneRanges[i].Lane().Compare(laneRanges[j].Lane()) < 0 })
		return &pb.Proposal{
			View:       ViewConv.Encode(m.view),
			Timestamp:  TimeConv.Encode(m.timestamp),
			LaneRanges: LaneRangeConv.EncodeSlice(laneRanges),
			App:        AppProposalConv.EncodeOpt(m.app),
		}
	},
	Decode: func(m *pb.Proposal) (*Proposal, error) {
		view, err := ViewConv.DecodeReq(m.View)
		if err != nil {
			return nil, fmt.Errorf("view: %w", err)
		}
		laneRanges, err := LaneRangeConv.DecodeSlice(m.LaneRanges)
		if err != nil {
			return nil, fmt.Errorf("laneRanges: %w", err)
		}
		timestamp, err := TimeConv.DecodeReq(m.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("timestamp: %w", err)
		}
		app, err := AppProposalConv.DecodeOpt(m.App)
		if err != nil {
			return nil, fmt.Errorf("appQC: %w", err)
		}
		return newProposal(
			view,
			timestamp,
			laneRanges,
			app,
		), nil
	},
}

// FullProposalConv is the protobuf converter for FullProposal.
var FullProposalConv = protoutils.Conv[*FullProposal, *pb.FullProposal]{
	Encode: func(m *FullProposal) *pb.FullProposal {
		laneQCs := make([]*LaneQC, 0, len(m.laneQCs))
		for _, qc := range m.laneQCs {
			laneQCs = append(laneQCs, qc)
		}
		return &pb.FullProposal{
			Proposal:  SignedMsgConv[*Proposal]().Encode(m.proposal),
			LaneQcs:   LaneQCConv.EncodeSlice(laneQCs),
			AppQc:     AppQCConv.EncodeOpt(m.appQC),
			TimeoutQc: TimeoutQCConv.EncodeOpt(m.timeoutQC),
		}
	},
	Decode: func(m *pb.FullProposal) (*FullProposal, error) {
		proposal, err := SignedMsgConv[*Proposal]().DecodeReq(m.Proposal)
		if err != nil {
			return nil, fmt.Errorf("proposal: %w", err)
		}
		laneQCs, err := LaneQCConv.DecodeSlice(m.LaneQcs)
		if err != nil {
			return nil, fmt.Errorf("laneQCs: %w", err)
		}
		laneQCsMap := map[LaneID]*LaneQC{}
		for _, qc := range laneQCs {
			laneQCsMap[qc.Header().Lane()] = qc
		}
		appQC, err := AppQCConv.DecodeOpt(m.AppQc)
		if err != nil {
			return nil, fmt.Errorf("appQC: %w", err)
		}
		timeoutQC, err := TimeoutQCConv.DecodeOpt(m.TimeoutQc)
		if err != nil {
			return nil, fmt.Errorf("timeoutQC: %w", err)
		}
		return &FullProposal{proposal: proposal, laneQCs: laneQCsMap, appQC: appQC, timeoutQC: timeoutQC}, nil
	},
}
